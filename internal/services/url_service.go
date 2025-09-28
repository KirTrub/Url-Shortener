package services

import (
	"context"

	"github.com/thanhpk/randstr"
)

type UrlService struct {
	r UrlRepo
}

func NewService(r UrlRepo) *UrlService {
	return &UrlService{r: r}
}

func (s *UrlService) AddNewLink(fullUrl string, c *context.Context) (string, error) {
	shortUrl, err := s.r.AddLink(randstr.String(8), fullUrl, c)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (s *UrlService) GetById(id string, c *context.Context) (string, error) {
	return s.r.GetById(id, c)
}
