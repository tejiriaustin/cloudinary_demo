package database

import (
	"cloudinary_demo/Thirdparty/excelservice"
	"cloudinary_demo/model"
)

func (r *repo) Save(data model.Data) error {
	err := excelservice.SaveToFile(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) FindAll() ([][]string, error) {
	t, err := excelservice.OpenAndRead()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *repo) Find() error {
	return nil
}
