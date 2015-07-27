package main

import (
	"flag"
	"time"

	"github.com/RobotClubKut/SlackBot/webapi"
	"github.com/RobotClubKut/slack/bot"
)

func main() {
	/*
		attachments := slack.NewAttachments()
		attachments.Attachments[0].Text = "にゃんぱす"
		js, _ := json.Marshal(attachments)
		fmt.Println(string(js))

		client := &http.Client{}
		data := url.Values{"payload": {string(js)}}
		resp, _ := client.Post(
			"https://hooks.slack.com/services/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",
			"application/x-www-form-urlencoded",
			strings.NewReader(data.Encode()),
		)
		ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
	*/
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
}
