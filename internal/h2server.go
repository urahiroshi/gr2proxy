package internal

import (
	"fmt"
	"net"
	"net/http"

	"github.com/urahiroshi/gr2proxy/config"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
)

type H2Server struct {
	server    http2.Server
	handler   http.HandlerFunc
	localPort int
	logger    *zap.Logger
}

func NewH2Server(c *config.Config) *H2Server {
	s := new(H2Server)
	s.localPort = c.LocalPort
	s.logger = c.Logger
	s.server = http2.Server{}
	return s
}

func (s *H2Server) Serve(handler http.HandlerFunc) {
	server := http2.Server{}
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.localPort))
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			s.logger.Fatal(err.Error())
		}
		server.ServeConn(conn, &http2.ServeConnOpts{Handler: http.Handler(handler)})
	}
}
