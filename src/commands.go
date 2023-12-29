package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func commandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Working!")
}

func commandImage(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	googleArgs := args[1:]
	googleString := strings.Join(googleArgs, "_")

	if len(googleString) < 13 {

		googleKey := viper.GetString("google.key")
		googleCx := viper.GetString("google.cx")

		googleRequest := ("https://www.googleapis.com/customsearch/v1?key=" + googleKey + "&searchType=image&safe=active&cx=" + googleCx + "&q=" + googleString)

		res, err := http.Get(googleRequest)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var response Response
		json.Unmarshal(body, &response)

		var imageLinkJSON = response.Response[0]

		s.ChannelMessageSend(m.ChannelID, imageLinkJSON.ImageLink)

	} else {
		s.ChannelMessageSend(m.ChannelID, "Search parameter cannot exceed 13 characters.")
	}
}

func commandGivePet(s *discordgo.Session, m *discordgo.MessageCreate) {
	givePetsMessage := ("Pets " + "<@" + m.Author.ID + ">")
	s.ChannelMessageSend(m.ChannelID, givePetsMessage)
}
