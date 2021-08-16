package managers

import (
	"cache/app/conf"
	"cache/app/datasources"
	"cache/objects"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type baseCacheSuit struct {
	suite.Suite
}

func (s *baseCacheSuit) SetupTest() {

}

func TestBaseCacheSuitInit(t *testing.T) {
	suite.Run(t, new(baseCacheSuit))
}

func (s *baseCacheSuit) TestBaseCache_Get() {

	s.Run("ObjectNotFound", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.OlderFistEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		cache := &baseCache{storage: storage}
		o, err := cache.Get("1")
		s.NotNil(err)
		s.Nil(o)
		s.Equal(objectNotFound, err)
	})

	s.Run("GetObjectWithInfinityTTL", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.OlderFistEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		storage.Add("1", objects.NewObject(`{"text": "one"}`))
		storage.Add("2", objects.NewObject(`{"text": "zwei"}`))
		storage.Add("3", objects.NewObject(`{"text": "trois"}`))

		cache := &baseCache{storage: storage}

		o, err := cache.Get("2")
		s.Nil(err)
		s.NotNil(o)
	})

	s.Run("GetExpiredObject", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.OlderFistEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		storage.Add("1", objects.NewObjectWithTTL(`{"text": "one"}`, 2))
		cache := &baseCache{storage: storage}

		time.Sleep(time.Second * 3)
		o, err := cache.Get("1")
		s.Nil(o)
		s.NotNil(err)
		s.Equal(objectNotFound, err)
	})

	s.Run("GetObjectWithTTLAlive", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.OlderFistEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		storage.Add("1", objects.NewObjectWithTTL(`{"text": "one"}`, 10))
		cache := &baseCache{storage: storage}

		time.Sleep(time.Second * 3)
		o, err := cache.Get("1")
		s.NotNil(o)
		s.Nil(err)
	})
}

func (s *baseCacheSuit) TestBaseCache_Delete() {

	s.Run("DeleteExistentObject", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.OlderFistEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		storage.Add("1", objects.NewObjectWithTTL(`{"text": "one"}`, 2))
		cache := &baseCache{storage: storage}

		o, err := cache.Delete("1")
		s.NotNil(o)
		s.Nil(err)
	})

	s.Run("DeleteNonExistentObject", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.OlderFistEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		cache := &baseCache{storage: storage}

		o, err := cache.Delete("1")
		s.Nil(o)
		s.NotNil(err)
		s.Equal(objectNotFound, err)
	})
}
