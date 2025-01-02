package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/yuki-miyakawa/random-team/internal/usecase"
)

func main() {
	godotenv.Load()
	e := echo.New()
	e.GET("/team", usecase.MakeTeam)
	e.GET("/weapon2", usecase.MainWeaponTwo)
	e.GET("/weapon3", usecase.MainWeaponThree)
	e.GET("/subweapon2", usecase.SubWeaponTwo)
	e.GET("/subweapon3", usecase.SubWeaponThree)
	e.GET("/specialweapon2", usecase.SpecialWeaponTwo)
	e.GET("/specialweapon3", usecase.SpecialWeaponThree)
	e.GET("/unified/special2", usecase.UnifiedSpecialWeaponTwo)
	e.GET("/unified/special3", usecase.UnifiedSpecialWeaponThree)
	e.GET("/unified/special2/versus", usecase.UnifiedSpecialWeaponVersusTwo)
	e.GET("/unified/special3/versus", usecase.UnifiedSpecialWeaponVersusThree)
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
