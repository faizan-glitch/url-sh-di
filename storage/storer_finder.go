package storage

import (
	"url-sh/models"
)

type StorerFinder interface {
	Store(*models.Url) error

	Find(string) (*models.Url, error)
}
