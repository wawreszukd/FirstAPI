package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}
func returnArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint returnSingleArticle hited")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", returnArticles)
	myRouter.HandleFunc("/article/{id}", returnArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article content"},
		Article{Id: "2", Title: "Title2", Desc: "Test", Content: "Test"},
	}

	handleRequest()
}
