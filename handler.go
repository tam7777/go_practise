package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Document(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getDocuments(w, r)
	case "POST":
		postDocument(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

type Article struct {
	Title string `json:"title"`
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ここにデータベースから情報を取得して返すコードを書く
	p := Article{
		Title: "title 1",
	}
	s, err := json.Marshal(p) //pは構造体、sはjson
	if err != nil {
		panic(err.Error())
	}
	w.Write(s)
}

func postDocument(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ここにデータベースに情報を登録するコードを書く
	p := Article{
		Title: "title 1",
	}

	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, _ = r.Body.Read(body)
	var jsonBody Article
	if err := json.Unmarshal(body[:length], &jsonBody); err != nil {
		panic(err)
	}
	if jsonBody == p {
		s, _ := json.Marshal(jsonBody)
		w.Write(s)
	} else {
		fmt.Println("error")
	}
}
