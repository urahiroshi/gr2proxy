package internal

import (
	"crypto/tls"
	"net"
	"net/http/httputil"
	"net/url"

	"golang.org/x/net/http2"
)

func NewProxy(remoteUrl string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(remoteUrl)
	if err != nil {
		return nil, err
	}
	dial := func(network, addr string, cfg *tls.Config) (net.Conn, error) {
		return net.Dial(network, addr)
	}
	transport := &http2.Transport{
		AllowHTTP: true,
		DialTLS:   dial,
	}
	p := httputil.NewSingleHostReverseProxy(u)
	p.Transport = transport
	return p, nil
}
