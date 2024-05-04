package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"strconv"

	"gopkg.in/yaml.v3"

	"github.com/bwmarrin/discordgo"
)

// Declare global variables
// General bot config
var BotConfig BotConfigStruct
// Default countNumber
var countNumber = 0

func main() {
	// Load config file and umarshal
	data, err := os.ReadFile("../config/config.yml")

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := yaml.Unmarshal(data, &BotConfig); err != nil {
		fmt.Println(err)
		return
	}

	// Load count file
	countNumber = loadCountFile()

	// Create Discord session
	token := BotConfig.DiscordConfig.Token

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		ErrorLog.Println("Error creating Discord session -", err)
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
		//fmt.Println("error opening connection,", err)
		ErrorLog.Println("Error opening Discord connection-", err)
		return
	}

	// Escape on CTRL-C
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	InfoLog.Println("Started Bob.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close.
	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set playing status
	s.UpdateGameStatus(0, "BAD GO CODE COMING THROUGH")
	InfoLog.Println("Updated GameStatus")
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
			commandTest(s, m)
		case "!image":
			commandImage(s, m, args)
		case "!givepet":
			commandGivePet(s, m)
		case "!count":
			commandCount(s, m, args)
		}
	}
}

func loadCountFile() int {
	countFileContent, error := os.ReadFile("../files/count.txt")

	if error != nil {
		ErrorLog.Println("Could not read file")
	}

	countFileContentStr := string(countFileContent)

	countFileContentInt, err := strconv.Atoi(countFileContentStr)

	// Check if file content is int
	if err != nil {
		ErrorLog.Println("File content loaded is not a number, defaulting to 0 -", err)
	} else {
		countNumber = countFileContentInt
		InfoLog.Println("Count file loaded, count is:", countFileContentStr)
	}

	return countNumber
}