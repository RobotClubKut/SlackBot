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
	postText := r.PostFormValue("text")
	fmt.Fprintf(w, "{\"text\": \""+postText+"\"}")
}
