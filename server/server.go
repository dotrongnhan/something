package server

import (
	"example.com/hello/controller"
	"fmt"
	"net/http"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8080")
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/sum", controller.Sum)

	http.HandleFunc("/minus", controller.Minus)

	http.HandleFunc("/multi", controller.Multi)

	http.HandleFunc("/div", controller.Div)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Error when running server")
	}
}
