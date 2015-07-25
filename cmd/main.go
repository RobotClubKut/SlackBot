package main

import (
	"flag"
	"time"

	"github.com/RobotClubKut/SlackBot/bot"
)

func main() {
	var i = flag.Bool("i", false, "初期化するかどうか")
	flag.Parse()
	if *i {
		bot.InitialDatabase()
	}
	for {
		bot.PostNoSubNews("")
		time.Sleep(7 * time.Minute)
	}
}
