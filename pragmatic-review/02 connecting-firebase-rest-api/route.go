package main

import (
	"connecting-firebase-rest-api/entity"
	"connecting-firebase-rest-api/repository"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting  the posts"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	var post entity.Post

	// const jsonStream = `{"Id": 11, "Title": "Title 1", "Text": "Text 1"}`

	fmt.Println("JSON => ", post)

	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		fmt.Println("Error decode => ", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the posts array"}`))
		return
	}

	post.Id = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)

}
