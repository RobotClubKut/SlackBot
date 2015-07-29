package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/RobotClubKut/SlackBot/lib/conf"
	"github.com/RobotClubKut/SlackBot/lib/log"
)

func main() {
	c := conf.NewCoufigure()
	js, err := json.Marshal(c)
	log.Terminate(err)
	err = ioutil.WriteFile("./bot-example.json", js, 0644)
	log.Terminate(err)
}
