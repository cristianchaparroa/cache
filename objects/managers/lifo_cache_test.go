package managers

import (
	"cache/app/conf"
	"cache/app/datasources"
	"cache/objects"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type lifoCacheSuite struct {
	suite.Suite
}

func (s *lifoCacheSuite) SetupTest() {

}

func TestLIFOCacheSuiteInit(t *testing.T) {
	suite.Run(t, new(lifoCacheSuite))
}

func (s *lifoCacheSuite) TestLIFOCache_Add() {

	s.Run("NonExceedingSlotLimit", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.OlderFistEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		cache := NewLIFOCache(storage)

		isAdded := cache.Add("1", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("2", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		oldestElement := storage.Front()
		oldestKey := oldestElement.Key.(string)

		expectedOldKey := "1"
		expectedCapacity := 5
		expectedLen := 2
		expectedIsFull := false

		s.Equal(expectedOldKey, oldestKey)
		s.Equal(expectedLen, storage.Len())
		s.Equal(expectedCapacity, storage.Capacity())
		s.Equal(expectedIsFull, storage.IsFull())
	})

	s.Run("ExceedingSlotLimit", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.NewestFirstEvictionPolicy,
		}
		storage := datasources.NewMemoryStorage(c)
		cache := NewLIFOCache(storage)

		isAdded := cache.Add("1", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("2", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("3", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("4", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("5", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("6", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		oldestElement := storage.Front()
		oldestKey := oldestElement.Key.(string)

		expectedOldestKey := "6"
		expectedCapacity := 5
		expectedLen := 5
		expectedIsFull := true

		s.Equal(expectedOldestKey, oldestKey)
		s.Equal(expectedLen, storage.Len())
		s.Equal(expectedCapacity, storage.Capacity())
		s.Equal(expectedIsFull, storage.IsFull())
	})

	s.Run("NonExpiredObjectWithTTL", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.NewestFirstEvictionPolicy,
		}
		storage := datasources.NewMemoryStorage(c)
		cache := NewLIFOCache(storage)

		isAdded := cache.Add("1", objects.NewObjectWithTTL(`{"text": "ok"}`, 10))
		s.Equal(true, isAdded)

		isAdded = cache.Add("2", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		time.Sleep(time.Second * 4) // sleep to simulate 4 seconds elapsed
		oldest, err := cache.Get("1")
		s.Nil(err)
		s.NotNil(oldest)
	})

	s.Run("ExpiredObjectWithTTL", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.NewestFirstEvictionPolicy,
		}
		storage := datasources.NewMemoryStorage(c)
		cache := NewLIFOCache(storage)

		isAdded := cache.Add("1", objects.NewObjectWithTTL(`{"text": "ok"}`, 10))
		s.Equal(true, isAdded)

		isAdded = cache.Add("2", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		time.Sleep(time.Second * 2) // sleep to simulate 2 seconds elapsed
		oldest, err := cache.Get("1")
		s.Nil(err)
		s.NotNil(oldest)

		time.Sleep(time.Second * 10) // sleep to simulate 10 seconds elapsed
		oldest, err = cache.Get("1")
		s.NotNil(err)
		s.Nil(oldest)
	})
}
