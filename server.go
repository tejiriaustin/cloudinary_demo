package main

import (
	server "cloudinary_demo/cmd/REST"
	"cloudinary_demo/controller"
	"fmt"
)

var (
	route server.Router = server.NewMuxRouter()
)

func main() {
	route.POST("/add", controller.AddPost)
	route.GET("/get", controller.GetPost)
	route.GET("/register", controller.Register)

	fmt.Println("Serving running on port 8000")
	route.SERVE(":8000")
}
