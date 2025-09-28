package api

import (
	"url-shortener/internal/models"
	"url-shortener/internal/services"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UrlHandler struct {
	c *context.Context
	s *services.UrlService
}

func NewUrlHandler(r *services.UrlService, c *context.Context) *UrlHandler {
	return &UrlHandler{s: r, c: c}
}

func (h *UrlHandler) NewLink(ctx *fiber.Ctx) error {
	request := new(models.Request)

	fullUrl := ctx.BodyParser(&request)
	if fullUrl != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
		return fmt.Errorf("cannot parse JSON: %v", fullUrl)
	}

	shortUrl, err := h.s.AddNewLink(request.FullUrl, h.c)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot create short link",
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"fullLink":  request.FullUrl,
		"shortLink": shortUrl,
	})
	return nil
}

func (h *UrlHandler) GetLink(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
		return fmt.Errorf("id is required")
	}

	fullUrl, err := h.s.GetById(id, h.c)
	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "link not found",
		})
		return err
	}

	ctx.Redirect(fullUrl, 302)
	return nil
}
