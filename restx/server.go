package restx

import (
    "github.com/shuitai/coney-framework/core/logx"
    "log"
)

type (
    Server struct {
        ngin *engine
    }
)

func MustNewServer(c RestConf) *Server {
    engine, err := NewServer(c)
    if err != nil {
        log.Fatal(err)
    }

    return engine
}

func NewServer(c RestConf) (*Server, error) {
    if err := c.SetUp(); err != nil {
        return nil, err
    }

    server := &Server{
        ngin: newEngine(c),
    }

    return server, nil
}

func (e *Server) AddRoutes(rs []Route) {
    e.ngin.AddRoutes(rs)
}

func (e *Server) AddRoute(r Route) {
    e.AddRoutes([]Route{r})
}

func (e *Server) Start() error{
    return e.ngin.Start()
}

func (e *Server) Stop() error {
    logx.Close()
    return e.ngin.Stop()
}

func (e *Server) Use(middleware Middleware) {
    e.ngin.Use(middleware)
}
