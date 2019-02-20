package db

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// New configures the database infrastructure
func New(postgres *gorm.DB, redis *redis.Client) *CacheRepo {
	return &CacheRepo{
		db:    &PostgresRepo{db: postgres},
		cache: redis,
	}
}
