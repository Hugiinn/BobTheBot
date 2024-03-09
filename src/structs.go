package main

// Main config
type BotConfigStruct struct {
	DiscordConfig DiscordConfigStruct `yaml:"discord"`
	GoogleConfig  GoogleConfigStruct  `yaml:"google"`
}

type DiscordConfigStruct struct {
	Token string `yaml:"token"`
}

type GoogleConfigStruct struct {
	Key string `yaml:"key"`
	Cx  string `yaml:"cx"`
}

// Commands
type Response struct {
	Response []items `json:"items"`
}

type items struct {
	ImageLink string `json:"link"`
}
