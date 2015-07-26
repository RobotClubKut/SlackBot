package webapi

import (
	"fmt"
	"net/http"
)

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
		if userName == "testUser" {
			fmt.Fprintf(w, "{\"text\": \""+text+"\"}")
		}
	}
}
