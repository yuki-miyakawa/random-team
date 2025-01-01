package usecase

import (
	"cmp"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/yuki-miyakawa/discord"
)

var (
	FourMembers = 4
	FiveMembers = 5
	SixMembers  = 6
	msg         = "team created"
)

type Member struct {
	Name   string
	Number int
}

func dividedIntoThreeAndOne(members []Member) error {
	teamText := fmt.Sprintf(
		"## ---Alpha Team---\n%s\n%s\n%s\n\n## ---Bravo Team---\n%s\n",
		members[0].Name,
		members[1].Name,
		members[2].Name,
		members[3].Name,
	)
	if err := discord.SendMessage(teamText); err != nil {
		return err
	}
	return nil
}

func dividedIntoFourAndOne(members []Member) error {
	teamText := fmt.Sprintf(
		"## ---Alpha Team---\n%s\n%s\n%s\n%s\n\n## ---Bravo Team---\n%s\n",
		members[0].Name,
		members[1].Name,
		members[2].Name,
		members[3].Name,
		members[4].Name,
	)
	if err := discord.SendMessage(teamText); err != nil {
		return err
	}
	return nil
}

func dividedIntoFourAndTwo(members []Member) error {
	teamText := fmt.Sprintf(
		"## ---Alpha Team---\n%s\n%s\n%s\n%s\n\n## ---Bravo Team---\n%s\n%s\n",
		members[0].Name,
		members[1].Name,
		members[2].Name,
		members[3].Name,
		members[4].Name,
		members[5].Name,
	)
	if err := discord.SendMessage(teamText); err != nil {
		return err
	}
	return nil
}

func MakeTeam(c echo.Context) error {
	godotenv.Load()
	memStr := os.Getenv("MEMBERS")
	memSlice := strings.Split(memStr, ",")
	members := make([]Member, len(memSlice))
	for i, v := range memSlice {
		members[i] = Member{
			Name:   v,
			Number: rand.Intn(100),
		}
	}
	slices.SortFunc(members, func(i, j Member) int {
		return cmp.Compare(i.Number, j.Number)
	})
	switch len(members) {
	case FourMembers:
		if err := dividedIntoThreeAndOne(members); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	case FiveMembers:
		if err := dividedIntoFourAndOne(members); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	case SixMembers:
		if err := dividedIntoFourAndTwo(members); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	default:
		return c.JSON(http.StatusBadRequest, "members must be 4, 5, or 6")
	}

	return c.JSON(http.StatusOK, msg)
}
