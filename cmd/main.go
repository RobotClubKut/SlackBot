package main

import (
	"fmt"

	"github.com/RobotClubKut/SlackBot/lib/nosub"
)

func main() {
	//_, err := ioutil.ReadFile("test.go")
	//log.TerminateAndWriteMessage(err, "hoge")
	fmt.Println(nosub.GetNosubUpdate())
}
