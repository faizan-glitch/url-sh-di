package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-sh/storage"
)

type urlController struct {
	store storage.StorerFinder
}

func NewUrlController(s storage.StorerFinder) *urlController {
	return &urlController{
		store: s,
	}
}

func (c *urlController) ShortenUrl(w http.ResponseWriter, r *http.Request) {

	body := ShortenRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON came"))
		return
	}

	c.store.Find(body.Url)
	// c.store.Find()

	// if myDb == "mongo" {
	// 	c.MongoStore.Store()
	// }
	// else if myDb == "memory" {
	// 	c.MemoryStore.Store()
	// }

	// c.MemoryStore.Find()
	// c.MongoStore.Find()
	// c.RedisStore.Find()
	// c.FileStore.Find()
	// c.DynamoStore.Find()

	fmt.Println(body.Url)
}

func (c *urlController) ResolveUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("resolve req came")
}

type ShortenRequest struct {
	Url       string `json: "url"`
	Password  string `json:"password"`
	IsPrivate bool   `json:"is_private"`
}
