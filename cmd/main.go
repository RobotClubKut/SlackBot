package main

import (
	"github.com/RobotClubKut/SlackBot/bot"
	"github.com/RobotClubKut/SlackBot/slack"
)

func main() {
	data := bot.GetNoSubData()
	slack.PostSlack(data)

	/*
		i := flag.Bool("i", false, "初期化するかどうか")
		t := flag.Int64("t", 7, "更新時間")

		flag.Parse()
		if *i {
			bot.InitialDatabase()
		}
		ch0 := make(chan bool)
		ch1 := make(chan bool)
		go func() {
			for {
				bot.PostNoSubNews("")
				time.Sleep(time.Duration(*t) * time.Minute)

			}
			ch0 <- true
		}()
		go func() {
			webapi.ViewWebApi()
			ch1 <- true
		}()
		<-ch0
		<-ch1
	*/
}
