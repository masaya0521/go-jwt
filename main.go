package main

import (
	"fmt"
	"net/http"

	"github.com/masaya0521/go-jwt/handler"
)

func main() {
	fmt.Print("test")

	http.HandleFunc("/", handler.HelloHandler)
	http.ListenAndServe(":8080", nil)
}