package db

import (
	"context"
	"net"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateDbClient() *redis.Client {
    host := os.Getenv("REDIS_HOST")
    pass := os.Getenv("REDIS_PASS")
    port := os.Getenv("REDIS_PORT")
    rdb := redis.NewClient(&redis.Options{
        Addr:     net.JoinHostPort(host, port),
        Password: pass, // no password set
        DB:       0,  // use default DB
    })

    return rdb
}
