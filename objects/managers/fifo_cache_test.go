package managers

import (
	"cache/app/conf"
	"cache/app/datasources"
	"cache/objects"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type fifoCacheSuit struct {
	suite.Suite
}

func (s *fifoCacheSuit) SetupTest() {

}

func TestFIFOCacheSuitInit(t *testing.T) {
	suite.Run(t, new(fifoCacheSuit))
}

func (s *fifoCacheSuit) TestFIFOCache_Add() {

	s.Run("NonExceedingSlotLimit", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.NewestFirstEvictionPolicy,
		}

		storage := datasources.NewMemoryStorage(c)
		cache := NewFIFOCache(storage)

		isAdded := cache.Add("1", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("2", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		newestElement := storage.Back()
		newestKey := newestElement.Key.(string)

		expectedLastKey := "2"
		expectedCapacity := 5
		expectedLen := 2
		expectedIsFull := false

		s.Equal(expectedLastKey, newestKey)
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
		cache := NewFIFOCache(storage)

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

		newestElement := storage.Back()
		newestKey := newestElement.Key.(string)

		expectedLastKey := "6"
		expectedCapacity := 5
		expectedLen := 5
		expectedIsFull := true

		s.Equal(expectedLastKey, newestKey)
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
		cache := NewFIFOCache(storage)

		isAdded := cache.Add("1", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("2", objects.NewObjectWithTTL(`{"text": "ok"}`, 10))
		s.Equal(true, isAdded)

		time.Sleep(time.Second * 4) // sleep to simulate 4 seconds elapsed
		newest, err := cache.Get("2")
		s.Nil(err)
		s.NotNil(newest)
	})

	s.Run("ExpiredObjectWithTTL", func() {
		c := &conf.Config{
			Slots:  5,
			TTL:    0,
			Policy: conf.NewestFirstEvictionPolicy,
		}
		storage := datasources.NewMemoryStorage(c)
		cache := NewFIFOCache(storage)

		isAdded := cache.Add("1", objects.NewObject(`{"text": "ok"}`))
		s.Equal(true, isAdded)

		isAdded = cache.Add("2", objects.NewObjectWithTTL(`{"text": "ok"}`, 5))
		s.Equal(true, isAdded)

		time.Sleep(time.Second * 2) // sleep to simulate 2 seconds elapsed
		newest, err := cache.Get("2")
		s.Nil(err)
		s.NotNil(newest)

		time.Sleep(time.Second * 10) // sleep to simulate 10 seconds elapsed
		newest, err = cache.Get("2")
		s.NotNil(err)
		s.Nil(newest)
	})
}
