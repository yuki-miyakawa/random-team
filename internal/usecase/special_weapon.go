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

var specialWeaponTwo = []string{
	"マルチミサイル",
	"ジェットパック",
	"ハイパープレッサー",
	"スーパーチャクチ",
	"インクアーマー",
	"ボムピッチャー",
	"アメフラシ",
	"イカスフィア",
	"バブルランチャー",
	"ナイスダマ",
	"ウルトラハンコ",
}

var specialWeaponThree = []string{
	"ウルトラショット",
	"グレートバリア",
	"ショクワンダー",
	"マルチミサイル",
	"アメフラシ",
	"ナイスダマ",
	"ホップソナー",
	"キューインキ",
	"メガホンレーザー5.1ch",
	"ジェットパック",
	"ウルトラハンコ",
	"カニタンク",
	"サメライド",
	"トリプルトルネード",
	"エナジースタンド",
	"デコイチラシ",
	"テイオウイカ",
	"ウルトラチャクチ",
	"スミナガシート",
	"スーパーチャクチ",
}

type MemberSpecialWeapon struct {
	Name          string
	SpecialWeapon string
}

func SpecialWeaponTwo(c echo.Context) error {
	memStr := os.Getenv("MEMBERS")
	memSlice := strings.Split(memStr, ",")
	members := make([]MemberSpecialWeapon, len(memSlice))
	for i, v := range memSlice {
		members[i] = MemberSpecialWeapon{
			Name:          v,
			SpecialWeapon: specialWeaponTwo[rand.Intn(len(specialWeaponTwo))],
		}
	}
	text := make([]string, len(members))
	for i, v := range members {
		text[i] = fmt.Sprintf("%s: %s", v.Name, v.SpecialWeapon)
	}
	joinedText := strings.Join(text, "\n")
	if err := discord.SendMessage(joinedText); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "random sub-weapon 2")
}

func SpecialWeaponThree(c echo.Context) error {
	memStr := os.Getenv("MEMBERS")
	memSlice := strings.Split(memStr, ",")
	members := make([]MemberSpecialWeapon, len(memSlice))
	for i, v := range memSlice {
		members[i] = MemberSpecialWeapon{
			Name:          v,
			SpecialWeapon: specialWeaponThree[rand.Intn(len(specialWeaponThree))],
		}
	}
	text := make([]string, len(members))
	for i, v := range members {
		text[i] = fmt.Sprintf("%s: %s", v.Name, v.SpecialWeapon)
	}
	joinedText := strings.Join(text, "\n")
	if err := discord.SendMessage(joinedText); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to send message")
	}
	return c.JSON(http.StatusOK, "random sub-weapon 3")
}
