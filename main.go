package main

import (
	"context"
	"log"
	"os"

	"UwU/handlers"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dep := handlers.Dependencies{}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatalln("DB_HOST required")
	}

	cacheHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		log.Fatalln("DB_HOST required")
	}

	app := fiber.New()

	// migration
	m, err := migrate.New(
		"file://schemas/migrations",
		dbHost)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatalln(err)
		return
	}
	m.Close()

	// connect to db
	db, err := pgxpool.Connect(context.Background(), dbHost)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer db.Close()
	dep.DB = db

	// connect to cache
	dep.Cache = redis.NewClient(&redis.Options{
		Addr:     cacheHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	handlers.InvalidateCache(context.TODO())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("deps", &dep)
		return c.Next()
	})

	app.Get("/articles", handlers.GetAllArticles)

	app.Get("/article/:id", handlers.GetArticle)

	app.Post("/article", handlers.CreateArticle)

	app.Listen(":3000")
}
