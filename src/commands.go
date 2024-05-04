package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"strconv"
	"os"

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
		}

		var response Response
		json.Unmarshal(body, &response)

		if len(response.Response) > 0 {
			jsonLength := len(response.Response)
			var randomNumber = rand.Intn(jsonLength - 0)

			var imageLinkJSON = response.Response[randomNumber]

			s.ChannelMessageSend(m.ChannelID, imageLinkJSON.ImageLink)
			InfoLog.Println(m.Author.Username, "(", m.Author.ID, ")", "used !image with args:", args)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Google returned an empty result, try another search parameter.")
			ErrorLog.Println(m.Author.Username, "(", m.Author.ID, ")", "used !image with args:", args, "but the response from Google is empty.")
		}

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

func commandCount(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	countArgs := args[1]

	countArgNumber, err := strconv.Atoi(countArgs)

	// Check if argument given was a number
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Argument is not a number.")
		ErrorLog.Println(m.Author.Username, "(", m.Author.ID, ")", "requested !count with an invalid argument.")
	} else {
		// Add one to countArgNumber to check it
		countNumberTmp := countNumber
		countNumberTmp += 1

		// Check if argument provided is the correct number
		if countNumberTmp == countArgNumber {
			countNumber += 1
			countNumberString := strconv.Itoa(countNumber)

			s.ChannelMessageSend(m.ChannelID, countNumberString)
			InfoLog.Println(m.Author.Username, "(", m.Author.ID, ")", "requested !count with", countArgNumber, "which was correct, adding one number to", countNumber, "to", countNumberTmp)

			countBytes := []byte(countNumberString)

			err := os.WriteFile("../files/count.txt", countBytes, 0644)

			if err != nil {
				ErrorLog.Println("Could not write counting number to file -", err)
			}

			InfoLog.Println("Wrote new number to file:", countNumber)

		} else {
			s.ChannelMessageSend(m.ChannelID, "That's not the correct number, idiot.")
			InfoLog.Println(m.Author.Username, "(", m.Author.ID, ")", "requested !count with", countArgNumber, "which was incorrect, answer should have been", countNumberTmp)
		}
	}
}