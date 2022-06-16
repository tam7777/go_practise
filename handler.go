package main

import (
	"encoding/json"
	"errors"
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
	err := json.NewEncoder(w).Encode(p) //pは構造体、sはjson, json.Marshal(p)でもいいけどAPIレスポンスの場合はjson.NewEncode(p).Marhsal -> string Encoder -> steam
	if err != nil {
		log.Fatal(err)
	}
}

func postDocument(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ここにデータベースに情報を登録するコードを書く
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, _ = r.Body.Read(body) //body 書く

	var jsonBody Article //受信した型の宣言
	if err := json.Unmarshal(body[:length], &jsonBody); err != nil {
		log.Fatal(err)
	} //jsonを構造体に変換

	var p Article
	p.Title = "title 1" //structにtitle 1を入れる
	if err := validation(jsonBody, p); err != nil {
		log.Fatal(err)
	} //受信した型と既にある型が同じか判別

	err := json.NewEncoder(w).Encode(jsonBody)
	if err != nil {
		log.Fatal(err)
	} //構造体をjsonに変換
}

func validation(jsonBody, p Article) error {
	if jsonBody != p {
		return errors.New("Error: %s") //errorのfuncはerrors.New()を使う。中にstring必要
	}
	return nil
}
