package cached

import (
	"fmt"
	"golang-starter/config"
	"log"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var once sync.Once

type RedisImpl struct {
	db    *redis.Client
	cache *cache.Cache
}

func NewRedisClient() *RedisImpl {
	log.Println("Initialize Redis connection")
	host := fmt.Sprintf("%s:%s", config.Get().Cache.Redis.Host, config.Get().Cache.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: config.Get().Cache.Redis.Pass, // no password set
		DB:       config.Get().Cache.Redis.DB,   // use default DB
	})

	ctx := rdb.Context()
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Redis Connection: ", err)
	}
	log.Println("Redis Connection: ", ping)

	cache := cache.New(&cache.Options{
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
		Redis:      rdb,
	})

	return &RedisImpl{
		db:    rdb,
		cache: cache,
	}
}

func (c RedisImpl) DB() *redis.Client {
	return c.db
}

func (c RedisImpl) Cache() *cache.Cache {
	return c.cache
}

func (c RedisImpl) Close() error {
	return c.db.Close()
}
