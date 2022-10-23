package handler

import (
	"encoding/json"
	"net/http"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "goで認証機能を試してみる",
		Tag:   "golang",
	}
	json.NewEncoder(w).Encode(post)
}