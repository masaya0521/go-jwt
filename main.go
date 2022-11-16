package main

import (
	"log"
	"net/http"

	"github.com/masaya0521/go-jwt/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/public", handler.PublicHandler)
	http.HandleFunc("/auth",  handler.GetTokenHandler)
	//http.HandleFunc("/private", handler.JwtMiddleware.Handler(handler.Private))

	if  err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}