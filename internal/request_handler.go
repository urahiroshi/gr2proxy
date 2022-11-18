package internal

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// RequestHandler object is not for a request, but for all requests
type RequestHandler struct {
	remoteUrl *url.URL
	logger    *zap.Logger
}

type InvalidURLError struct{}

func (err *InvalidURLError) Error() string {
	return "Invalid URL"
}

func NewRequestHandler(remoteUrl string, logger *zap.Logger) (*RequestHandler, error) {
	u, err := url.Parse(remoteUrl)
	if err != nil {
		return nil, err
	}
	return &RequestHandler{
		remoteUrl: u,
		logger:    logger,
	}, nil
}

func (rh *RequestHandler) GetNames(r *http.Request) (string, string, error) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) < 3 {
		rh.logger.Warn("Invalid URL",
			zap.String("url", r.URL.Path),
		)
		return "", "", new(InvalidURLError)
	}
	serviceName := paths[1]
	methodName := paths[2]
	return serviceName, methodName, nil
}

func (rh *RequestHandler) newReflectionClient(ctx context.Context) (*grpcreflect.Client, error) {
	refConn, err := grpc.Dial(rh.remoteUrl.Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	stub := rpb.NewServerReflectionClient(refConn)
	return grpcreflect.NewClient(ctx, stub), nil
}

func (rh *RequestHandler) GetMethodDescription(r *http.Request, ctx context.Context) (*desc.MethodDescriptor, error) {
	serviceName, methodName, err := rh.GetNames(r)
	if err != nil {
		return nil, err
	}
	rh.logger.Debug("get service name and method name",
		zap.String("serviceName", serviceName),
		zap.String("methodName", methodName),
	)
	// TODO: use cache or something
	client, err := rh.newReflectionClient(ctx)
	if err != nil {
		return nil, err
	}
	serviceDesc, err := client.ResolveService(serviceName)
	if err != nil {
		return nil, err
	}
	methodDesc := serviceDesc.FindMethodByName(methodName)
	return methodDesc, nil
}

func (rh *RequestHandler) ReadBodyAsJson(r *http.Request, methodDesc *desc.MethodDescriptor) ([]byte, error) {
	grpcBytes, err := ioutil.ReadAll(r.Body)
	defer func() {
		r.Body.Close()
		// recover body
		r.Body = ioutil.NopCloser(bytes.NewReader(grpcBytes))
	}()
	if err != nil {
		return nil, err
	}
	protoBytes := grpcBytes[grpcHeaderLength:]
	messageDesc := methodDesc.GetInputType()
	return toJson(protoBytes, messageDesc)
}
