package main

import (
	"fmt"

	"github.com/RobotClubKut/SlackBot/lib/mysql"
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

	fmt.Println(mysql.GetAnimeMostNewAnime())
}
