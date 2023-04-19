package database

import "github.com/go-redis/redis"

var Cl *redis.Client

func SetUpRedis() {
	Cl = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}