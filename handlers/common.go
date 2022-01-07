package handlers

import (
	pb "UwU/proto"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Dependencies struct {
	DB    *pgxpool.Pool
	Cache *redis.Client
	pb.UnimplementedUwUServer
}
