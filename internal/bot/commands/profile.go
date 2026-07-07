package commands

import (
	"dota/app/internal/opendota"
	"fmt"
	"math"

	"github.com/bwmarrin/discordgo"
)

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
	return &discordgo.MessageEmbed{
		Title: profile.Profile.Personaname,
		Color: 0x00BFFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: profile.Profile.Avatarfull,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "ПОТАНЦИАЛЬНЫЙ MMR",
				Value:  fmt.Sprintf("%v", math.Round(profile.ComputedMmr)),
				Inline: true,
			},
			{
				Name:  "SteamId",
				Value: fmt.Sprintf("%v", profile.Profile.Steamid),
			},
		},
	}
}
