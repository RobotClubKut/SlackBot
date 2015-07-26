package main

import (
	"flag"
	"time"

	"github.com/RobotClubKut/SlackBot/bot"
	"github.com/RobotClubKut/SlackBot/webapi"
)

func main() {
	i := flag.Bool("i", false, "初期化するかどうか")
	t := flag.Int64("t", 7, "更新時間")
	flag.Parse()
	if *i {
		bot.InitialDatabase()
	}
	ch0 := make(chan bool)
	ch1 := make(chan bool)
	for {
		go func() {
			bot.PostNoSubNews("")
			time.Sleep(time.Duration(*t) * time.Minute)
			ch0 <- true
		}()
		go func() {
			webapi.ViewWebApi()
			ch1 <- true
		}()
		<-ch0
		<-ch1
	}
}
