package main

import (
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	sa := option.WithCredentialsFile("D:/Program/Golang/Tutorial/learn-go-e8508-firebase-adminsdk-gc2xs-885278405d.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	post := getPost()
	log.Print(post)

	result, err := client.Collection("post").Doc("eec0NLIITjVUE49FXsNy").Set(context.Background(), post)
	if err != nil {
		log.Fatalln(err)
		log.Print("test2")
	}

	log.Print(result)

	defer client.Close()
}

func getPost() Post {
	return Post{
		Id:    9,
		Title: "Title 9",
		Text:  "Text 9",
	}
}
