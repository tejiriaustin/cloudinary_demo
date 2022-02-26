package database

import (
	"net/http"
	"os"
)

func getfile() error {
	fileAddress := os.Getenv("CLOUDINARY_URL")

	resp, err := http.Get(fileAddress)
	if err != nil {
		return http.ErrMissingFile
	}
	defer resp.Body.Close()

	return nil
}
func (r *repo) Save(int642 int64) error {
	return nil
}

func (r *repo) FindAll() error {
	err := getfile()
	if err != nil {
		return http.ErrMissingFile
	}

	return nil
}

func (r *repo) Find() error {
	return nil
}
