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
	fmt.Println("command flag:", strings.Contains(text, "deny:"))
	fmt.Println("token check:", configure.OutgoingSlackConf.Token == token)
	fmt.Println("user name:", configure.OutgoingSlackConf.UserName != userName)
	if strings.Contains(text, "deny:") && configure.OutgoingSlackConf.Token == token && configure.OutgoingSlackConf.UserName != userName {
		words := strings.Split(text, "deny:")
		var bufStr string
		for _, w := range words {
			if w != "" {
				bufStr += w
			}
		}
		words = strings.Split(bufStr, " ")
		bufStr = ""
		for _, w := range words {
			if w != "" {
				bufStr += w
			}
		}
		words = strings.Split(bufStr, ":")
		var bufStrings []string
		for _, w := range words {
			if w != "" {
				bufStrings = append(bufStrings, w)
			}
		}
		words = bufStrings

		fmt.Println("words:", words)
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
		if configure.OutgoingSlackConf.UserName == userName {

		} else {
			postText := "{\"text\":\"nilぱすー\"}"
			fmt.Fprintf(w, postText)
		}
	}
}

//ViewWebPage is view web api
func ViewWebPage() {
	configure := conf.ReadConfigure()
	http.HandleFunc("/", home)
	http.HandleFunc("/deny", deny)
	http.ListenAndServe(":"+configure.OutgoingSlackConf.Port, nil)
}
