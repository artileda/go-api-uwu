package handlers_test

import (
	"UwU/handlers"
	"context"
	"log"
	"os"
	"testing"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis"
	"github.com/golang-migrate/migrate"
	"github.com/jackc/pgx/v4/pgxpool"
)

func connection() *handlers.Dependencies {
	return nil
}

func TestArticle(t *testing.T) {
	c := context.TODO()
	dep := connection()
	newArticle := handlers.ArticleDTO{}
	param := handlers.ArticleParams{}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatalln("DB_HOST required")
	}

	// migration
	m, err := migrate.New(
		"file://schemas/migrations",
		dbHost)
	if err != nil {
		t.Fatal(err)
		return
	}
	m.Up()
	m.Close()

	// connect to db
	db, err := pgxpool.Connect(context.Background(), dbHost)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer db.Close()
	dep.DB = db

	// connect to cache
	dep.Cache = cache.New(&cache.Options{
		Redis: redis.NewRing(&redis.RingOptions{
			Addrs: map[string]string{
				"localhost": ":6379",
			},
		}),
	})

	err = handlers.PersistArticle(c, *dep, newArticle)
	if err != nil {
		t.Fatal(err)
		return
	}

	result := handlers.FetchArticlesFromDB(c, *dep, param)
	if len(result) == 0 {
		t.Fatalf("Artikel gagal masuk")
		return
	} else {
		res, err := handlers.FetchArticle(c, *dep, result[0].ID)
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
