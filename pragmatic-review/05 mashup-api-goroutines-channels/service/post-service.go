package service

import (
	"errors"
	"mashup-api-goroutines-channels/entity"
	"mashup-api-goroutines-channels/repository"
	"math/rand"
)

var (
	repoService repository.PostRepository
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

func NewPostService(repo repository.PostRepository) PostService {
	repoService = repo
	return &service{}
}

func (*service) Validate(post *entity.Post) error {

	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The title is empty")
		return err
	}

	return nil

}

func (*service) Create(post *entity.Post) (*entity.Post, error) {

	post.Id = rand.Int63()
	return repoService.Save(post)

}

func (*service) FindAll() ([]entity.Post, error) {

	return repoService.FindAll()

}
