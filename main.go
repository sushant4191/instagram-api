package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}
type Users []user

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		user{ID: "Test ID", Name: "Test Name", Pass: "Test Pass"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Users)
}

func handleRequests() {
	http.HandleFunc("/", home)
	//	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/users", allUsers)
	// http.HandleFunc("/users/id", ids)
	// http.HandleFunc("/posts", posts)
	// http.HandleFunc("/posts/id", postid)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}
