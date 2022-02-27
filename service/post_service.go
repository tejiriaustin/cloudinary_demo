package service

import (
	"cloudinary_demo/database"
	"cloudinary_demo/model"
	validation "github.com/go-ozzo/ozzo-validation"
)

//TODO: INCLUDE VALIDATION LIBRARY

type post struct {
	Id int64 `json:"id"`
}

type Service interface {
	AddService(data model.Data) error
	FindService() error
	FindAllService() ([][]string, error)
	ValidateService() error
}

func NewService() Service {
	return &post{}
}

var storage = database.NewCloudStore()

func (n *post) AddService(data model.Data) error {
	err := storage.Save(data)
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

func (n *post) FindAllService() ([][]string, error) {
	t, err := storage.FindAll()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (n *post) ValidateService() error {
	return validation.ValidateStruct(&post{},
		validation.Field(n.Id, validation.Required, validation.NotNil))
}
