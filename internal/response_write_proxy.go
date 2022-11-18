package internal

import (
	"net/http"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"go.uber.org/zap"
)

type ResponseWriteProxy struct {
	writer     http.ResponseWriter
	methodDesc *desc.MethodDescriptor
	logger     *zap.Logger
	response   *Response
	http.ResponseWriter
}

func NewResponseWriteProxy(
	writer http.ResponseWriter,
	methodDesc *desc.MethodDescriptor,
	logger *zap.Logger,
	response *Response,
) *ResponseWriteProxy {
	wp := new(ResponseWriteProxy)
	wp.writer = writer
	wp.methodDesc = methodDesc
	wp.logger = logger
	wp.response = response
	return wp
}

func (wp *ResponseWriteProxy) Header() http.Header {
	return wp.writer.Header()
}

func (wp *ResponseWriteProxy) Write(b []byte) (int, error) {
	if len(b) > grpcHeaderLength {
		protoBytes := b[grpcHeaderLength:]
		messageDesc := wp.methodDesc.GetOutputType()
		jsonBytes, err := toJson(protoBytes, messageDesc)
		jsonStr := string(jsonBytes)
		if err == nil {
			wp.logger.Debug("Write() called",
				zap.String("body", jsonStr),
			)
			wp.response.Body = jsonStr
		} else {
			wp.logger.Warn(err.Error())
		}
	}
	return wp.writer.Write(b)
}

func (wp *ResponseWriteProxy) WriteHeader(statusCode int) {
	wp.logger.Debug("WriteHeader() called",
		zap.Int("statusCode", statusCode),
	)
	wp.response.StatusCode = statusCode
	wp.writer.WriteHeader(statusCode)
}

func (wp *ResponseWriteProxy) SaveHeaders() {
	for name, values := range wp.Header() {
		wp.logger.Debug("Headers",
			zap.String("name", name),
			zap.Strings("values", values),
		)
		// TODO: more cool way
		trailerPrefix := "Trailer:"
		if strings.Index(name, trailerPrefix) == 0 {
			wp.response.Trailer[strings.Replace(name, trailerPrefix, "", 1)] = values
		} else {
			wp.response.Header[name] = values
		}
	}
}
