package main

import (
	"time"

	"github.com/RobotClubKut/SlackBot/lib/mysql"
	"github.com/RobotClubKut/SlackBot/lib/slack"
)

func main() {
	//_, err := ioutil.ReadFile("test.go")
	//log.TerminateAndWriteMessage(err, "hoge")

	//noSub := nosub.GetNosubUpdate()
	//mysql.InsertNoSubData(noSub)
	//js, _ := json.Marshal(n)
	//fmt.Println(string(js))
	//nosub.GetAnimeData("slack_bot")
	//couchbase.InsertDataBase(noSub)

	//var word []string
	//word = append(word, "test")
	//word = append(word, "にゃんぱす")
	//mysql.InsertDenyWord(word)

	//fmt.Println(mysql.GetAnimeMostNewAnime())
	//a := slack.NewAttachments(1)
	//a.Attachments[0].Text = "test"
	//js, _ := json.Marshal(a)
	//slack.Post(string(js))
	//mysql.InsertNoSubBufData(noSub)
	//fmt.Println(mysql.DiffNoSubData("[720p]だんちがい 第04話 「四号棟 カレーライス」"))
	/*
		for _, d := range mysql.CreatePostNoSubData() {
			fmt.Println(d)
		}
	*/
	ch := make(chan bool)
	go func() {
		for {
			slack.PostAnimeInfomation(mysql.CreatePostNoSubData())
			time.Sleep(10 * time.Minute)
		}

		ch <- true
	}()
	<-ch
}
