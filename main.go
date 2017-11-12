package main

import (
	"github.com/kotamat/slack"
	"github.com/k0kubun/pp"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		pp.Print(err)
		return
	}

	channel := os.Getenv("channel")
	token := os.Getenv("token")
	cmd := os.Getenv("cmd")
	text := os.Getenv("text")

	api := slack.New(token)
	option := slack.MsgOptionCommand(cmd, text)
	res, timestamp, text, err := api.SendMessage(channel, option)
	pp.Println(res, timestamp, text, err)
}
