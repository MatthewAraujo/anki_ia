package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/MatthewAraujo/anki_ia/cmd/api"
	configs "github.com/MatthewAraujo/anki_ia/config"
	database "github.com/MatthewAraujo/anki_ia/db"
	"github.com/MatthewAraujo/anki_ia/pkg/assert"
	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/utils"
	"github.com/redis/go-redis/v9"
)

var logger = utils.NewParentLogger("starting API")

func main() {
	url := configs.Envs.Postgres.URL

	redisCfg := redis.Options{
		DB:       configs.Envs.Redis.Database,
		Password: configs.Envs.Redis.Password,
		Addr:     fmt.Sprintf("%s:%s", configs.Envs.Redis.Address, configs.Envs.Redis.Port),
	}

	db, err := database.NewMyPostgresSQLStorage(url)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	redis := database.NewRedisStorage(redisCfg)

	err = healthRedis(redis)
	assert.NoError(err, "Redis is offline")
	logger.Info("Connect to redis")

	repository := repository.New(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.API.Port), repository, db, redis)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("DB: Successfully connected!")
}

func healthRedis(redisClient *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
