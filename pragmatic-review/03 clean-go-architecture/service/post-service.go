package service

import (
	"clean-go-architecture/entity"
	"clean-go-architecture/repository"
	"errors"
	"math/rand"
)

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

func NewPostService() PostService {
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
	return repo.Save(post)

}

func (*service) FindAll() ([]entity.Post, error) {

	return repo.FindAll()

}
