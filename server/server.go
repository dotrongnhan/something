package server

import (
	"example.com/hello/math"
	"fmt"
	"net/http"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8080")
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/sum", math.Sum)

	http.HandleFunc("/minus", math.Minus)

	http.HandleFunc("/multi", math.Multi)

	http.HandleFunc("/div", math.Div)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Error when running server")
	}
}
