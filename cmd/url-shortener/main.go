package main

import (
	"context"
	"url-shortener/internal/api"
	"url-shortener/internal/repo"
	"url-shortener/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	db := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		DB:       0,
		Password: "",
	})

	_, err := db.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repo.New(db)
	service := services.NewService(repo)
	handler := api.NewUrlHandler(service, &ctx)

	app := fiber.New()

	app.Post("/new", handler.NewLink)
	app.Get("/:id", handler.GetLink)

	app.Listen(":8082")

}
