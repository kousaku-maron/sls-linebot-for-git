package services

import (
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	channelID     string = os.Getenv("CHANNEL_ID")
	channelSecret string = os.Getenv("CHANNEL_SECRET")
)

func initBot() (*linebot.Client, error) {
	bot, err := linebot.New(channelID, channelSecret)

	if err != nil {
		return nil, err
	}

	return bot, nil
}
