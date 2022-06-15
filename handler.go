package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func getDocuments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ここにデータベースから情報を取得して返すコードを書く
	var p Article
	p.Title = "title 1"
	err := json.NewEncoder(w).Encode(p) //pは構造体、sはjson, json.Marshal(p)でもいいけどAPIレスポンスの場合はjson.NewEncode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func postDocument(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ここにデータベースに情報を登録するコードを書く
	var p Article
	p.Title = "title 1"

	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, _ = r.Body.Read(body)
	var jsonBody Article

	if err := json.Unmarshal(body[:length], &jsonBody); err != nil {
		log.Fatal(err)
	}

	if jsonBody == p {
		err := json.NewEncoder(w).Encode(jsonBody)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Did not get title:title 1")
	}
}

/*
func validation(jsonBody, p Article) error {
	if jsonBody == p {

	} else {
		fmt.Println("Did not get title:title 1")
	}
}
*/
