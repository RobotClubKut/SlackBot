package webapi

// webapiだお

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RobotClubKut/SlackBot/lib/conf"
	"github.com/RobotClubKut/SlackBot/lib/mysql"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func deny(w http.ResponseWriter, r *http.Request) {
	configure := conf.ReadConfigure()
	text := r.PostFormValue("text")
	token := r.PostFormValue("token")
	userName := r.PostFormValue("user_name")
	fmt.Println("test")
	if strings.Contains(text, "deny:") && configure.OutgoingSlackConf.Token == token && configure.OutgoingSlackConf.UserName == userName {
		fmt.Println("catch")
		text = strings.Replace(text, "deny:", "", 0)
		text = strings.Replace(text, " ", "", 0)
		words := strings.Split(text, ":")
		mysql.InsertDenyWord(words)
		//postText := "{\"text\":\"" +  + ""\"}"
		postText := "{\"text\":\""
		for _, w := range words {
			if w != "" {
				postText += w
				postText += ","
			}
		}
		postText += "\"}"
		fmt.Fprintf(w, postText)
	} else {
		postText := "{\"text\":\"nilぱすー\"}"
		fmt.Fprintf(w, postText)
	}
}

//ViewWebPage is view web api
func ViewWebPage() {
	configure := conf.ReadConfigure()
	http.HandleFunc("/", home)
	http.ListenAndServe(":"+configure.OutgoingSlackConf.Port, nil)
}
