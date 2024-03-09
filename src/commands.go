package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func commandTest(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Working!")
	InfoLog.Println(m.Author.Username, "(", m.Author.ID, ")", "used !test")
}

func commandImage(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// Declare BotConfig for Google key/cx
	googleArgs := args[1:]
	googleString := strings.Join(googleArgs, "_")

	if len(googleString) < 13 {

		googleKey := BotConfig.GoogleConfig.Key
		googleCx := BotConfig.GoogleConfig.Cx

		googleRequest := ("https://www.googleapis.com/customsearch/v1?key=" + googleKey + "&searchType=image&safe=active&cx=" + googleCx + "&q=" + googleString)

		res, err := http.Get(googleRequest)
		if err != nil {
			ErrorLog.Println(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			ErrorLog.Println(err)
		}

		var response Response
		json.Unmarshal(body, &response)

		var imageLinkJSON = response.Response[0]

		s.ChannelMessageSend(m.ChannelID, imageLinkJSON.ImageLink)
		InfoLog.Println(m.Author.Username, "(", m.Author.ID, ")", "used !image with args:", args)

	} else {
		s.ChannelMessageSend(m.ChannelID, "Search parameter cannot exceed 13 characters.")
		InfoLog.Println(m.Author.Username, "(", m.Author.ID, ")", "requested !image with args longer than 13 characters.")
	}
}

func commandGivePet(s *discordgo.Session, m *discordgo.MessageCreate) {
	givePetsMessage := ("Pets " + "<@" + m.Author.ID + ">")
	s.ChannelMessageSend(m.ChannelID, givePetsMessage)
	InfoLog.Println(m.Author.Username, "(", m.Author.ID, ")", "used !givepets")
}
