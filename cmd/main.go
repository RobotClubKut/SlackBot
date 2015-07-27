package main

import (
	"flag"
	"time"

	"github.com/RobotClubKut/SlackBot/bot"
	"github.com/RobotClubKut/SlackBot/slack"
	"github.com/RobotClubKut/SlackBot/webapi"
)

func main() {

	i := flag.Bool("i", false, "初期化するかどうか")
	t := flag.Int64("t", 7, "更新時間")
	p := flag.String("p", "8080", "port")

	flag.Parse()
	if *i {
		bot.InitialDatabase()
	}
	ch0 := make(chan bool)
	ch1 := make(chan bool)
	go func() {
		for {
			//bot.PostNoSubNews("")
			data := bot.GetNoSubData()
			slack.PostSlack(data)

			time.Sleep(time.Duration(*t) * time.Minute)

		}
		ch0 <- true
	}()
	go func() {
		webapi.ViewWebApi(*p)
		ch1 <- true
	}()
	<-ch0
	<-ch1
}
