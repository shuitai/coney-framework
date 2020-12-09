package httpx

import (
    "github.com/gofiber/fiber/v2"
    "github.com/shuitai/coney-framework/core/logx"
    "net/http"
)

func Error(c *fiber.Ctx, code int, err error) {
    c.Status(code)
    logx.Error(err.Error())
}

func ErrorCode(c *fiber.Ctx, code int) {
    c.Status(code)
}

func ErrorJson(c *fiber.Ctx, code int, v interface{}, err error) {
    c.Status(code)
    c.JSON(v)
    logx.Error(err.Error())
}

func Ok(c *fiber.Ctx) {
    c.Status(http.StatusOK)
}

func OkJson(c *fiber.Ctx, v interface{}) {
    c.Status(http.StatusOK)
    c.JSON(v)
}