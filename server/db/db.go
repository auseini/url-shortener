package db

import (
    "context"
    "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateDbClient() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    return rdb
}
