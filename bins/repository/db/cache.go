package db

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
	"github.com/hugocortes/hooks-api/bins"
	"github.com/hugocortes/hooks-api/bins/models"
	"github.com/hugocortes/hooks-api/common/cache"
	gModels "github.com/hugocortes/hooks-api/models"
)

// CacheRepo ..
type CacheRepo struct {
	DB    bins.DB
	Cache *redis.Client
}

// GetAll ...
func (r *CacheRepo) GetAll(accountID string, opts *gModels.QueryOpts) ([]*models.Bin, error) {
	return r.DB.GetAll(accountID, opts)
}

// Get ...
func (r *CacheRepo) Get(accountID string, ID string) (*models.Bin, error) {
	cacheKey := cache.GenKey("Get", accountID, ID)
	cachedValue := r.Cache.Get(cacheKey)
	bin := &models.Bin{}

	var err error
	if cachedValue.Val() != "" {
		err := json.Unmarshal([]byte(cachedValue.Val()), bin)
		if err != nil {
			log.Printf("unmarshalling caching error: %s", err.Error())
			return nil, err
		}
	} else {
		bin, err = r.DB.Get(accountID, ID)
		if err != nil {
			return nil, err
		}

		marshalled, err := json.Marshal(bin)
		if err != nil {
			log.Printf("marshalling caching error: %s", err.Error())
			return nil, err
		}

		r.Cache.Set(cacheKey, marshalled, cache.Expiration())
	}

	return bin, err
}

// Create ...
func (r *CacheRepo) Create(bin *models.Bin) (string, error) {
	return r.DB.Create(bin)
}

// Update ...
func (r *CacheRepo) Update(accountID string, ID string, bin *models.Bin) error {
	r.Cache.Del(cache.GenKey("Get", accountID, ID))

	return r.DB.Update(accountID, ID, bin)
}

// Delete ...
func (r *CacheRepo) Delete(accountID string, ID string) error {
	r.Cache.Del(cache.GenKey("Get", accountID, ID))

	return r.DB.Delete(accountID, ID)
}

// Destroy ...
func (r *CacheRepo) Destroy(accountID string) error {
	return r.DB.Destroy(accountID)
}
