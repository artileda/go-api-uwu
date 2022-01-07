package handlers_test

import (
	"UwU/handlers"
	"context"
	"log"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestArticle(t *testing.T) {
	c := context.TODO()
	dep := handlers.Dependencies{}
	newArticle := handlers.ArticleDTO{
		AuthorID: 1,
		Title:    "Chicken running by its feet",
		Body:     "Chicken running by its feet, but the feet its tasty for humans.",
	}
	// Anya is authors with id 1
	param := handlers.ArticleParams{
		Author: "Anya",
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatalln("DB_HOST required")
	}

	cacheHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		log.Fatalln("DB_HOST required")
	}

	// migration
	m, err := migrate.New(
		"file://../schemas/migrations",
		dbHost)
	if err != nil {
		t.Fatal(err)
		return
	}
	m.Up()
	defer m.Down()

	// connect to db
	db, err := pgxpool.Connect(context.Background(), dbHost)
	if err != nil {
		t.Fatal(err)
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

	err = handlers.PersistArticle(c, dep, newArticle)
	if err != nil {
		t.Fatal(err)
		return
	}

	result, err := handlers.FetchArticlesFromDB(c, dep, param)
	if err != nil {
		t.Fatal(err)
		return
	}
	if len(result) == 0 {
		t.Fatalf("Artikel gagal masuk")
		return
	} else {
		res, err := handlers.FetchArticle(c, dep, result[0].ID)
		if err != nil {
			t.Fatal(err)
			return
		}
		if res == nil {
			t.Fatal("Artikel tidak ditemukan")
			return
		}
	}
}
