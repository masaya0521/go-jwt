package handler

import (
	"encoding/json"
	"net/http"
)

type postPrivate struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
}

var Private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
		Tag:   "Go",
	}
	json.NewEncoder(w).Encode(post)
})