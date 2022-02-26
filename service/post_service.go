package service

import (
	"cloudinary_demo/database"
	"math/rand"

	validation "github.com/go-ozzo/ozzo-validation"
)

//TODO: INCLUDE VALIDATION LIBRARY

type post struct {
	Id int64 `json:"id"`
}

type Service interface {
	AddService() error
	FindService() error
	FindAllService() error
	ValidateService() error
}

func NewService() Service {
	return &post{}
}

var (
	storage database.Storage = database.NewCloudStore()
)

func (n *post) AddService() error {
	Id := rand.Int63()
	err := storage.Save(Id)
	if err != nil {
		return err
	}
	return nil
}

func (n *post) FindService() error {
	err := storage.Find()
	if err != nil {
		return nil
	}
	return nil
}

func (n *post) FindAllService() error {
	err := storage.FindAll()
	if err != nil {
		return err
	}
	return nil
}

func (n *post) ValidateService() error {
	return validation.ValidateStruct(&post{},
		validation.Field(n.Id, validation.Required, validation.NotNil))
}
