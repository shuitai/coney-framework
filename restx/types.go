package restx

import (
    "github.com/gofiber/fiber/v2"
)

type
(
    Middleware func(c *fiber.Ctx) error

    Route struct {
        Method       string
        Path         string
        RouteHandler fiber.Handler
    }
)
