package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author *User `json:"author"`
}

var posts []Post = []Post{}

func main() {
	posts = append(posts,Post{Title:"post 1",Body:"This is my first post",Author:&User{FullName:"sharukh khan",UserName:"sharukh010",Email:"sharukh010@gmail.com"}})

	posts = append(posts,Post{Title:"post 2",Body:"This is my second post",Author:&User{FullName:"sharukh khan",UserName:"sharukh010",Email:"sharukh010@gmail.com"}})

	router := mux.NewRouter()

	router.HandleFunc("/posts",getPosts).Methods("GET")

	fmt.Println("Server starting on post 8081")
	
	if err := http.ListenAndServe(":8081",router);err!=nil{
		log.Fatalf("Error: %s",err.Error())
	}
}

func getPosts(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","Application/json")

	json.NewEncoder(w).Encode(posts)
}