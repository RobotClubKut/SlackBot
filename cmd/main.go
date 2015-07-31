package main

import (
	"time"

	"github.com/RobotClubKut/SlackBot/lib/mysql"
	"github.com/RobotClubKut/SlackBot/lib/slack"
	"github.com/RobotClubKut/SlackBot/lib/webapi"
)

func main() {
	chSlack := make(chan bool)
	chWebAPI := make(chan bool)
	go func() {
		for {
			slack.PostAnimeInfomation(mysql.CreatePostNoSubData())
			time.Sleep(10 * time.Minute)
		}

		chSlack <- true
	}()
	go func() {
		webapi.ViewWebPage()
		chWebAPI <- true
	}()
	<-chSlack
	<-chWebAPI
}
