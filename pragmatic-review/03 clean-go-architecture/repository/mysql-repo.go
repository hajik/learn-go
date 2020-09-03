package repository

import (
	"cloud.google.com/go/firestore"
	"connecting-firebase-rest-api/entity"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
)

type repo struct{}

//NewFirestoreRepository create a new repo
func NewMysqlRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "learn-go-e8508"
	CollectionName string = "post"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		fmt.Println("Failed to create a firestore client : ", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(CollectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		fmt.Println("Failed adding a new post : ", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		fmt.Println("Failed to create a firestore client : ", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	it := client.Collection(CollectionName).Documents(ctx)
	for {

		doc, err := it.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			fmt.Println("Failed to create a Documents list : ", err)
			return nil, err
		}

		post := entity.Post{
			Id:    doc.Data()["Id"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)

	}

	return posts, nil
}
