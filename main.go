package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/masaya0521/go-jwt/handler"
)

func main() {
	fmt.Print("test")

	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/public", handler.PublicHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}