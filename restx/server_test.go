package restx

import (
    "github.com/go-resty/resty/v2"
    "github.com/gofiber/fiber/v2"
    "github.com/shuitai/coney-framework/core/conf"
    "github.com/shuitai/coney-framework/core/logx"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestNewServer(t *testing.T) {
    _, err := NewServer(RestConf{})
    assert.NotNil(t, err)
}

func TestServerWithRoute(t *testing.T) {
    config_str := `
        Name: dashboard-api
        Host: 0.0.0.0
        Port: 8888
        
        Logger:
          Mode: console
          Level: info
   `
    var restConf RestConf
    conf.LoadConfigFromYamlBytes([]byte(config_str), &restConf)
    server, err := NewServer(restConf)

    result := "Hello, World 👋!"
    var Routes = [] Route {
        {
            "get", "/first/kevin/2017", func(ctx *fiber.Ctx) error {
                logx.Infof("header: %s", ctx.Context().Request.Header.Host())
                return ctx.SendString(result)
            },
        },
    }

    server.AddRoutes(Routes)
    assert.Nil(t, err)

    go func() {
        time.Sleep(time.Second * 5)
        url := "http://localhost:8888/coney/first/kevin/2017"
        client := resty.New()

        resp, err := client.R().
            EnableTrace().
            Get(url)
        assert.Nil(t, err)

        logx.Infof("response: %s", resp.String())

        assert.Equal(t,  resp.String(), result)

        time.Sleep(time.Second * 5)
        logx.Info("stopping server...")
        server.Stop()
    }()

    err = server.Start()
    assert.Nil(t, err)


}
