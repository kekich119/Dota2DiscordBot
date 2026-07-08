package commands

import (
	opendota "dota/app/internal/opendota"
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/bwmarrin/discordgo"
)

var heroes opendota.Heroes

func CommandShowProfile(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options

	accountID := options[0].StringValue()

	profile := opendota.SearchByAccountId(accountID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				buildProfileEmbed(profile),
			},
		},
	})
}

func buildProfileEmbed(profile opendota.Dotka) *discordgo.MessageEmbed {
	winrate := opendota.GetWinrateByAccountId(fmt.Sprint(profile.Profile.AccountID))
	procentWin := float64(winrate.Win) / float64(winrate.Win+winrate.Lose) * 100

	return &discordgo.MessageEmbed{
		Title: profile.Profile.Personaname,
		Color: 0x00BFFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: profile.Profile.Avatarfull,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "ПОТАНЦЕВАЛЬНЫЙ MMR",
				Value:  fmt.Sprintf("%v", math.Round(profile.ComputedMmr)),
				Inline: true,
			},
			{
				Name:  "Steam",
				Value: fmt.Sprintf("https://steamcommunity.com/profiles/%v", profile.Profile.Steamid),
			},
			{
				Name:   "ПОБЕДЫ",
				Value:  fmt.Sprintf("🟢 Win: %v", winrate.Win),
				Inline: true,
			},
			{
				Name:   "ПОРОЖЕНИЯ",
				Value:  fmt.Sprintf("🔴 Lose: %v", winrate.Lose),
				Inline: true,
			},
			{
				Name:   "ПРОЦЕНТ ПОБЕД",
				Value:  fmt.Sprintf("%.2f%%", math.Round(procentWin)),
				Inline: true,
			},
		},
	}
}

func getNameHeroesById(id int) (heroName string) {

	file, err := os.ReadFile("heroes.json")

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(file, &heroName)

	return heroName

}
