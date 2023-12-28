package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func commandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Working!")
}

func commandImage(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 2 {

		if len(args[1]) < 13 {
			googleKey := viper.GetString("google.key")
			googleCx := viper.GetString("google.cx")

			googleRequest := ("https://www.googleapis.com/customsearch/v1?key=" + googleKey + "&searchType=image&cx=" + googleCx + "&q=" + args[1])

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
			s.ChannelMessageSend(m.ChannelID, "Parameter too lang, doofus.")
		}

	} else if len(args) > 2 {
		s.ChannelMessageSend(m.ChannelID, "Too many parameters, doofus.")
	} else {
		s.ChannelMessageSend(m.ChannelID, "That requires a parameter, doofus.")
	}
}
