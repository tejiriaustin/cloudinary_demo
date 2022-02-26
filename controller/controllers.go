package controller

import (
	"cloudinary_demo/middleware"
	"cloudinary_demo/model"
	"cloudinary_demo/service"
	"encoding/json"
	"net/http"
)

var (
	post service.Service = service.NewService()
	req  model.Data
	m    middleware.Imiddleware = middleware.NewMiddleWare()
)

func Register(w http.ResponseWriter, r *http.Request) {
	token, err := m.CreateToken()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(token)

}

func GetPost(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(req)
	err = m.VerifyToken()
	if err != nil {
		return
	}

	if err = post.ValidateService(); err != nil {
		return
	}

	if err = post.FindAllService(); err != nil {
		return
	}
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return
	}

	if err = post.ValidateService(); err != nil {
		return
	}

	if err = post.AddService(); err != nil {
		return
	}
}
