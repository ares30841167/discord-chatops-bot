package main

import (
	"guanyu.dev/chatopsbot/internal/chatopsbot"
	"guanyu.dev/chatopsbot/internal/chatopsbot/util/progutil"
)

func main() {
	bot, err := chatopsbot.New()
	progutil.CheckErrorOccurred(err)

	_, err = bot.Start()
	progutil.CheckErrorOccurred(err)
}
