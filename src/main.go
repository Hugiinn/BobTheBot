package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

type Response struct {
	Response []Items `json:"items"`
}

type Items struct {
	ImageLink string `json:"link"`
}

func main() {

	viper.AddConfigPath("./config/")
	viper.SetConfigName("config.yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Create Discord session
	token := viper.GetString("bot.token")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register ready and messagecreate events
	dg.AddHandler(messageCreate)
	dg.AddHandler(ready)

	// Declare intents
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	// Open Discord & listen
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Escape on CTRL-C
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close.
	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set playing status
	s.UpdateGameStatus(0, "BAD GO CODE COMING THROUGH")
}

// Called whenever a message is sent
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore self
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Catch any responses starting with ! particularly sticker responses caused Bob to crash
	if strings.HasPrefix(m.Content, "!") {

		// Split response
		args := strings.Fields(m.Content)

		commandName := args[0]

		switch commandName {
		case "!test":
			commandHelp(s, m)
		case "!image":
			commandImage(s, m, args)
		case "!givepet":
			commandGivePet(s, m)
		}
	}
}
