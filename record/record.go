package record

import (
	"context"
	"net/http"
	"time"

	"github.com/urahiroshi/gr2proxy/config"
	"github.com/urahiroshi/gr2proxy/internal"
)

type RecordConfig struct {
	config.Config
	RemoteUrl string
}

func Run(c *RecordConfig) {
	p, err := internal.NewProxy(c.RemoteUrl)
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	s := internal.NewH2Server(&c.Config)
	rh, err := internal.NewRequestHandler(c.RemoteUrl, c.Logger)
	if err != nil {
		c.Logger.Fatal(err.Error())
	}
	fixtureMap := make(internal.FixtureMap)

	adminServer := internal.NewAdminServer(c.Logger, c.AdminPort, fixtureMap)
	go adminServer.Serve()

	s.Serve(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		methodDesc, err := rh.GetMethodDescription(r, ctx)
		if err != nil {
			c.Logger.Warn(err.Error())
			p.ServeHTTP(w, r)
			return
		}
		jsonBytes, nil := rh.ReadBodyAsJson(r, methodDesc)
		if err != nil {
			c.Logger.Warn(err.Error())
			p.ServeHTTP(w, r)
			return
		}
		jsonStr := string(jsonBytes)
		c.Logger.Info(jsonStr)

		response := internal.NewResponse()
		wp := internal.NewResponseWriteProxy(w, methodDesc, c.Logger, response)
		p.ServeHTTP(wp, r)
		wp.SaveHeaders()
		serviceName, methodName, err := rh.GetNames(r)
		if err != nil {
			c.Logger.Error(err.Error())
		} else {
			// TODO: now http header is ignored, but it should be handled with regular expression in future
			fixtureMap.SaveResponse(serviceName, methodName, "*", jsonStr, response)
			fixtureJson, _ := fixtureMap.ToJson()
			c.Logger.Info(string(fixtureJson))
		}
	})
}
