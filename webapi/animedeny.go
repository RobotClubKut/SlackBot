package webapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/RobotClubKut/SlackBot/slack"
)

type Conf struct {
	Token    string
	UserName string
}

func CreateConfExample() {
	c := Conf{Token: "", UserName: "slackbot"}
	js, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("../conf/webapi_configure-example.json", js, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadConfigure() *Conf {
	var ret Conf

	js, err := ioutil.ReadFile("../conf/webapi_configure.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(js, &ret)

	if err != nil {
		log.Fatalln(err)
	}
	return &ret
}

func ViewWebApi() {
	http.HandleFunc("/animedeny", animedeny)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func animedeny(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("text")
	token := r.PostFormValue("token")
	userName := r.PostFormValue("user_name")
	configure := ReadConfigure()

	if token == configure.Token {
		if userName != configure.UserName {
			attachments := slack.NewAttachments()
			attachments.Attachments[0].Text = text
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

			fmt.Fprintf(w, "")

			/*
				fmt.Println("outgoing-webhook: " + configure.UserName)
				fmt.Println("userName: " + userName)
				postString := "衝撃の事実. "
				postString += text
				postString += "受理できない."
				fmt.Fprintf(w, "{\"text\": \""+postString+"\"}")
			*/
		} else {
			fmt.Fprintf(w, "token does not match.")
		}
	}
}

//curl -F 'payload={"channel": "#bot-test", "username": "webhookbot", "text": "This is posted to #bot-test and comes from a bot named webhookbot.", "icon_emoji": ":ghost:"}' -F 'attachments={"attachments":[{"fallback":"","color":"#36a64f","pretext":"","author_name":"","author_link":"","author_icon":"","title":"","title_link":"","text":"衝撃の事実. にゃんぱすなのん受理できない.","fields":[{"title":"","value":"","short":false}],"image_url":"","thumb_url":""}]}' https://hooks.slack.com/services/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
//curl -F 'payload={"attachments":[{"fallback":"","color":"#000000","pretext":"","author_name":"","author_link":"","author_icon":"","title":"","title_link":"","text":"衝撃の事実. にゃんぱすなのん受理できない.","fields":[{"title":"","value":"","short":false}],"image_url":"","thumb_url":""}]}' https://hooks.slack.com/services/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
