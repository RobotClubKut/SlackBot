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

	if token == "xEwE8oq4UUm0pMrUsx3bgPGo" {
		if userName == "shirase_aoi" {
			fmt.Fprintf(w, "{\"text\": \""+text+"\"}")
		}
	}
}
