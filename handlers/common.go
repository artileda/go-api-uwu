package handlers

import (
	"github.com/go-redis/cache/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Dependencies struct {
	DB    *pgxpool.Pool
	Cache *cache.Cache
}
