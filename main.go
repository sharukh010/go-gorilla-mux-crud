package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	/*-----------------Adding data to the posts-----------*/
	posts = append(posts,Post{Title:"post 1",Body:"This is my first post",Author:&User{FullName:"sharukh khan",UserName:"sharukh010",Email:"sharukh010@gmail.com"}})

	posts = append(posts,Post{Title:"post 2",Body:"This is my second post",Author:&User{FullName:"sharukh khan",UserName:"sharukh010",Email:"sharukh010@gmail.com"}})

	/*--------------------------Routes--------------------*/
	router := mux.NewRouter()

	router.HandleFunc("/posts",getPosts).Methods("GET")
	router.HandleFunc("/posts",addPost).Methods("POST")
	router.HandleFunc("/posts/{id}",getPost).Methods("GET")
	router.HandleFunc("/posts/{id}",updatePost).Methods("PUT")

	/*-----------------------Starting the server-------------*/
	fmt.Println("Server starting on post 8081")

	if err := http.ListenAndServe(":8081",router);err!=nil{
		log.Fatalf("Error: %s",err.Error())
	}
}

func getPosts(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","Application/json")

	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","Application/json")

	var newPost Post 

	json.NewDecoder(r).Decode(&newPost)

	posts = append(posts,newPost)

	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)

	id,err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return 
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No data found with specified ID"))
		return 
	}

	post := posts[id]

	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","Application/json")

	var idParam string = mux.Vars(r)["id"]

	id,err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Id could not converted to Integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No data founded with specified ID"))
		return 
	}

	var updatedItem Post 

	json.NewDecoder(r.Body).Decode(&updatedItem)

	posts[id] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)
}