package storage

import (
	"errors"
	"url-sh/models"
)

type memoryStore struct {
	Data []*models.Url
}

func NewMemoryStore() *memoryStore {
	return &memoryStore{
		Data: []*models.Url{},
	}
}

func (store *memoryStore) Store(url *models.Url) error {

	existingUrl, err := store.Find(url.OriginalUrl)

	if err != nil {
		return err
	}

	if existingUrl != nil {
		return errors.New("this url already exists")
	}

	store.Data = append(store.Data, url)

	return nil
}

func (store *memoryStore) Find(longUrl string) (*models.Url, error) {
	for _, v := range store.Data {
		if v.OriginalUrl == longUrl {
			return v, nil
		}
	}

	return nil, nil
}
