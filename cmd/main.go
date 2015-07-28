package main

import (
	"io/ioutil"

	"github.com/RobotClubKut/SlackBot/lib/log"
)

func main() {
	_, err := ioutil.ReadFile("test.go")
	log.WriteErrorLogAndMessage(err, "test")
}
