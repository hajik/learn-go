package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Post struct {
	Id 		int 	`json:"id"`
	Title 	string 	`json:"title"`
	Text 	string 	`json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post Post

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&post); err != nil {
		fmt.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the posts array"}`))
		return
	}

	post.Id = len(posts)+1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(post)
	resp.Write(result)

}