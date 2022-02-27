package main

import (
	server "cloudinary_demo/cmd/REST"
	"cloudinary_demo/controller"
	"fmt"
	"os"
)

var (
	route = server.NewMuxRouter()
)

func main() {
	route.POST("/add", controller.AddPost)
	route.GET("/get", controller.GetPost)
	route.GET("/register", controller.Register)

	port := os.Getenv("PORT")
	fmt.Println(port)
	fmt.Println("Serving running on port 8000")
	route.SERVE(port)
}
