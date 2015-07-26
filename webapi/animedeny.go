package webapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Conf struct {
	Token    string
	UserName string
}

func CreateConfExample() {
	c := Conf{Token: "", UserName: ""}
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
			postString := "衝撃の事実. \n"
			postString += text
			postString += "受理できない.\n"
			fmt.Fprintf(w, "{\"text\": \""+postString+"\"}")
		}
	}
}
