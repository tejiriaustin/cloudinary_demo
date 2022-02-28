package controller

import (
	"cloudinary_demo/middleware"
	"cloudinary_demo/model"
	"cloudinary_demo/service"
	"encoding/json"
	"log"
	"net/http"
)

var (
	post service.Service = service.NewService()
	req  model.Data
	m    = middleware.NewMiddleWare()
)

func Register(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return
	}

	token, err := m.CreateToken()
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		return
	}

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

	t, err := post.FindAllService()
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(t)
}

func AddPost(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return
	}

	if err = post.ValidateService(); err != nil {
		return
	}

	new := model.Data{}
	err = json.NewDecoder(r.Body).Decode(new)
	if err != nil {
		log.Fatal(err)
	}

	if err = post.AddService(new); err != nil {
		return
	}
}
