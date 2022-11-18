package internal

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type handler struct {
	adminServer *AdminServer
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.adminServer.serveHTTP(w, r)
}

type AdminServer struct {
	fixtureMap FixtureMap
	logger     *zap.Logger
	server     *http.Server
}

func NewAdminServer(logger *zap.Logger, adminPort int, fixtureMap FixtureMap) *AdminServer {
	s := new(AdminServer)
	s.fixtureMap = fixtureMap
	s.logger = logger
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", adminPort),
		Handler: &handler{adminServer: s},
	}
	return s
}

func (s *AdminServer) HandleGetFixtures(w http.ResponseWriter, r *http.Request) {
	fixtureJson, err := s.fixtureMap.ToJson()
	if err != nil {
		s.logger.Error(err.Error())
		w.Write([]byte(`{"error": "internal error"}`))
		w.WriteHeader(500)
		return
	}
	w.Write(fixtureJson)
	w.WriteHeader(200)
}

func (s *AdminServer) serveHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "application/json")
	switch path {
	case "/fixtures":
		s.HandleGetFixtures(w, r)
	default:
		w.Write([]byte(`{"error": "not found"}`))
		w.WriteHeader(404)
	}
}

func (s *AdminServer) Serve() error {
	return s.server.ListenAndServe()
}
