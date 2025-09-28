package services

import "context"

type UrlRepo interface {
	AddLink(shortUrl, fullUrl string, c *context.Context) (string, error)
	GetById(id string, c *context.Context) (string, error)
}