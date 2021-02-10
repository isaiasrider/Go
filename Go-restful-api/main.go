package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	article := Articles{
		Article{Title: "Izarilhas",Desc: "Test Description", Content: "Hello Wolrd"},
	}

	fmt.Fprintf(w,"Endpoint Hit: All Kinda Books\n")
	json.NewEncoder(w).Encode(article)
}
func homePage(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "HomePage Endpoint Hit\t teste \n teste")
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081",nil))

}

func main() {
	handleRequests()
}