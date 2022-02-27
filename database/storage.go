package database

import "cloudinary_demo/model"

type repo struct {
	model.Data
}

//TODO: FLESH OUT THE DATABASE TO SUPPORT CLOUDINARY

type Storage interface {
	Save(data model.Data) error
	FindAll() ([][]string, error)
	Find() error
}

func NewCloudStore() Storage {
	return &repo{}
}
