package models

import "time"

type Url struct {
	Id           string    `json:"id"`
	OriginalUrl  string    `json:"original"`
	ShortenedUrl string    `json:"shortened"`
	ShortenCode  string    `json:"code"`
	Password     string    `json:"password"`
	IsPrivate    bool      `json:"is_private"`
	CreatedAt    time.Time `json:"created_at"`
}
