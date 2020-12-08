package restx

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "os"
)

type engine struct {
    conf                 RestConf
    middlewares          []fiber.Handler
    routes               []Route
    application          *fiber.App
    quitSignal           chan os.Signal
}

func newEngine(c RestConf) *engine {
    engine := &engine{
        conf: c,
        quitSignal :make(chan os.Signal),
    }

    return engine
}

func (s *engine) AddRoutes(r []Route) {
    s.routes = append(s.routes, r...)
}

func (s *engine) Use(middleware Middleware) {
    s.middlewares = append(s.middlewares, middleware)
}

func (s *engine) Start() error {
    s.application = fiber.New()
    s.application.Use(logger.New(logger.Config{
        TimeFormat: "2006-01-02T15:04:05",
    }))
    contextPath := s.conf.ContextPath
    group := s.application.Group(contextPath)
    for _, r := range s.routes {
        group.Add(r.Method, r.Path, r.RouteHandler)
    }

    return s.application.Listen(fmt.Sprintf(":%d", s.conf.Port))
}

func (s *engine) Stop() error {
    return s.application.Shutdown()
}

