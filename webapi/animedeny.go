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

type DenyList struct {
	List []string
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

func ViewWebApi(port string) {
	http.HandleFunc("/animedeny", animedeny)
	http.HandleFunc("/", home)
	http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func addDenyDB(str []string) {
	r, err := ioutil.ReadFile("../database/deny.db")
	if err != nil {
		log.Println(err)
	}
	var readList DenyList
	err = json.Unmarshal(r, &readList)
	if err != nil {
		log.Println("db is empty.")
		log.Println(err)
	}
	readList.List = append(readList.List, str...)

	w, err := json.Marshal(readList)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile("../database/deny.db", w, 0644)
}

func animedeny(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("text")
	token := r.PostFormValue("token")
	userName := r.PostFormValue("user_name")
	configure := ReadConfigure()

	if token == configure.Token {
		if userName != configure.UserName {
			confJs, err := ioutil.ReadFile("../conf/incoming_webhooks_configure.json")
			if err != nil {
				log.Fatalln(err)
			}
			var conf slack.Conf
			json.Unmarshal(confJs, &conf)

			attachments := slack.NewAttachments(1)
			if strings.HasPrefix(text, "deny:") {
				//text = "true: "
				str := strings.Trim(strings.TrimSpace(text), "deny:")
				fmt.Println("str: " + str)
				strList := strings.Split(str, ",")

				addDenyDB(strList)
			} else {
				text = "false"
			}
			attachments.Attachments[0].Text = text
			js, _ := json.Marshal(attachments)
			//fmt.Println(string(js))

			client := &http.Client{}
			data := url.Values{"payload": {string(js)}}
			resp, _ := client.Post(
				conf.ApiURL,
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
