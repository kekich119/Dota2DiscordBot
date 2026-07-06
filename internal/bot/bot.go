package bot

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

func Start() {

	file, _ := os.ReadFile("token.json")

	json.Unmarshal(file, token)

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	channelID := "1523759908409966795"
	guildID := "1523759906920992938"

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "!join" {
		connectToVoice(guildID, channelID, false, false, s)

	}

}

func connectToVoice(guildId string, channelId string, mute bool, deaf bool, s *discordgo.Session) {
	vc, err := s.ChannelVoiceJoin(guildId, channelId, mute, deaf)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	audio, err := os.ReadFile("01 - Семнадцать ножевых.mp3")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	vc.OpusSend <- audio

	fmt.Println("Bot connected into channel: ", vc.ChannelID)

}
