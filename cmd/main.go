package main

import "github.com/RobotClubKut/SlackBot/lib/nosub"

func main() {
	//_, err := ioutil.ReadFile("test.go")
	//log.TerminateAndWriteMessage(err, "hoge")
	//n := nosub.GetNosubUpdate()
	//js, _ := json.Marshal(n)
	//fmt.Println(string(js))
	nosub.GetAnimeData()
}
