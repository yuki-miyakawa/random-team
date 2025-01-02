package usecase

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yuki-miyakawa/discord"
)

var subWeaponTwo = []string{
	"スプラッシュボム",
	"キューバンボム",
	"クイックボム",
	"カーリングボム",
	"ロボットボム",
	"タンサンボム",
	"トーピード",
	"トラップ",
	"ポイズンミスト",
	"ポイントセンサー",
	"スプラッシュシールド",
	"ジャンプビーコン",
	"スプリンクラー",
}

var subWeaponThree = []string{
	"スプラッシュボム",
	"キューバンボム",
	"クイックボム",
	"スプリンクラー",
	"スプラッシュシールド",
	"タンサンボム",
	"カーリングボム",
	"ロボットボム",
	"ジャンプビーコン",
	"ポイントセンサー",
	"トラップ",
	"ポイズンミスト",
	"ラインマーカー",
	"トーピード",
}

type MemberSubWeapon struct {
	Name      string
	SubWeapon string
}

func SubWeaponTwo(c echo.Context) error {
	memStr := os.Getenv("MEMBERS")
	memSlice := strings.Split(memStr, ",")
	members := make([]MemberSubWeapon, len(memSlice))
	for i, v := range memSlice {
		members[i] = MemberSubWeapon{
			Name:      v,
			SubWeapon: subWeaponTwo[rand.Intn(len(subWeaponTwo))],
		}
	}
	text := make([]string, len(members))
	for i, v := range members {
		text[i] = fmt.Sprintf("%s: %s", v.Name, v.SubWeapon)
	}
	joinedText := strings.Join(text, "\n")
	if err := discord.SendMessage(joinedText); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "random sub-weapon 2")
}

func SubWeaponThree(c echo.Context) error {
	memStr := os.Getenv("MEMBERS")
	memSlice := strings.Split(memStr, ",")
	members := make([]MemberSubWeapon, len(memSlice))
	for i, v := range memSlice {
		members[i] = MemberSubWeapon{
			Name:      v,
			SubWeapon: subWeaponThree[rand.Intn(len(subWeaponThree))],
		}
	}
	text := make([]string, len(members))
	for i, v := range members {
		text[i] = fmt.Sprintf("%s: %s", v.Name, v.SubWeapon)
	}
	joinedText := strings.Join(text, "\n")
	if err := discord.SendMessage(joinedText); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "random sub-weapon 3")
}
