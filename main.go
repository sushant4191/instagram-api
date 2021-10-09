package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	ID    string `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
type Users []user
type post struct {
	ID        string `json:"ID"`
	Caption   string `json:"caption"`
	URL       string `json:"img_url"`
	Timestamp string `json:"timestamp"`
}
type Posts []post

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		user{ID: "Test ID", Name: "Test Name", Email: "Email", Pass: "Test Pass"},
	}

	fmt.Println("Endpoint Hit: All Posts Endpoint")
	json.NewEncoder(w).Encode(users)
}
func allPosts(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		post{ID: "Test ID", Caption: "Test CAP", URL: "testurl", Timestamp: "time"},
	}
	fmt.Println("Endpoint Hit: All Posts Endpoint")
	json.NewEncoder(w).Encode(posts)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", home)
	//	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/users", allUsers)
	// http.HandleFunc("/users/id", ids)
	http.HandleFunc("/posts", allPosts)
	// http.HandleFunc("/posts/id", postid)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}
