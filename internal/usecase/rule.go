package usecase

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuki-miyakawa/discord"
)

func UnifiedSpecialWeaponTwo(c echo.Context) error {
	special := specialWeaponTwo[rand.Intn(len(specialWeaponTwo))]
	if err := discord.SendMessage(special); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "unified special-weapon 2")
}

func UnifiedSpecialWeaponThree(c echo.Context) error {
	special := specialWeaponThree[rand.Intn(len(specialWeaponThree))]
	if err := discord.SendMessage(special); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "unified special-weapon 3")
}

func UnifiedSpecialWeaponVersusTwo(c echo.Context) error {
	specials := make([]string, 2)
	for i := range specials {
		specials[i] = specialWeaponTwo[rand.Intn(len(specialWeaponTwo))]
	}
	text := fmt.Sprintf("## ---Alpha Team---\n%s\n## ---Bravo Team---\n%s\n", specials[0], specials[1])
	if err := discord.SendMessage(text); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "random sub-weapon 2")
}

func UnifiedSpecialWeaponVersusThree(c echo.Context) error {
	specials := make([]string, 2)
	for i := range specials {
		specials[i] = specialWeaponThree[rand.Intn(len(specialWeaponThree))]
	}
	text := fmt.Sprintf("## ---Alpha Team---\n%s\n## ---Bravo Team---\n%s\n", specials[0], specials[1])
	if err := discord.SendMessage(text); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "random sub-weapon 2")
}
