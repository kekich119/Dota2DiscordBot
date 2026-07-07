package bot

import (
	"dota/app/internal/opendota"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type token struct {
	Token string `json:"token"`
}

var commands []*discordgo.ApplicationCommand

func Start() {

	var Token token
	file, err := os.ReadFile("internal/bot/token.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &Token)

	discord, err := discordgo.New("Bot " + Token.Token)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	discord.AddHandler(interactionCreate)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		fmt.Print(err)
		return
	}

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "profile",
			Description: "Показать профиль игрока",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "account_id",
					Description: "ID Аккаунта в доте",
					Required:    true,
				},
			},
		},
		{
			Name:        "hello",
			Description: "hello",
		},
	}

	for _, cmd := range commands {
		_, err := discord.ApplicationCommandCreate(
			discord.State.User.ID,
			"",
			cmd,
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	switch i.ApplicationCommandData().Name {

	case "profile":
		options := i.ApplicationCommandData().Options

		accountID := options[0].StringValue()

		profile := opendota.SearchByAccountId(accountID)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%f", profile.ComputedMmr),
			},
		})

	}
}
