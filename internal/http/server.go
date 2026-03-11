package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo"

	"github.com/mustan989/discord-sandbot/internal/config"
)

func Serve(ctx context.Context, config config.HTTP) error {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	e.POST("/health", func(c echo.Context) error {
		message := c.QueryParam("message")
		if message == "" {
			message = "allo"
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": message,
		})
	})

	address := fmt.Sprintf(":%d", config.Port)

	slog.InfoContext(ctx, "HTTP server is starting", "address", address)

	return e.Start(address)
}
