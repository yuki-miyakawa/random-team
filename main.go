package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuki-miyakawa/random-team/internal/usecase"
)

func main() {
	e := echo.New()
	e.GET("/random", usecase.MakeTeam)
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
