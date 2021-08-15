package datasources

import (
	"cache/app/conf"
	"github.com/stretchr/testify/suite"
	"testing"
)

type memoryStorageSuite struct {
	suite.Suite
}

func (s *memoryStorageSuite) SetupTest() {

}

func TestMemoryStorageSuitInit(t *testing.T) {
	suite.Run(t, new(memoryStorageSuite))
}

func (s *memoryStorageSuite) TestMemory_capacity() {
	s.Run("MemoryInitCapacity", func() {
		c := &conf.Config{Slots: 10}

		expectedLen := 0
		expectedCapacity := 10

		memory := NewMemoryStorage(c)
		s.Equal(memory.Len(), expectedLen)
		s.Equal(memory.Capacity(), expectedCapacity)
	})
}

func (s *memoryStorageSuite) TestMemory_Add() {
	s.Run("MemoryAddRepeatedKeys", func() {
		c := &conf.Config{Slots: 10}
		memory := NewMemoryStorage(c)
		memory.Add("1", `{"text": "ok"}`)
		memory.Add("1", `{"text": "unit testing"}`)
		memory.Add("1", `{"text": "cb insights"}`)

		expectedLen := 1
		expectedCapacity := 10

		s.Equal(memory.Capacity(), expectedCapacity)
		s.Equal(memory.Len(), expectedLen)
	})

	s.Run("MemoryAddSimple", func() {
		c := &conf.Config{Slots: 10}
		memory := NewMemoryStorage(c)
		memory.Add("1", `{"text": "ok"}`)
		memory.Add("2", `{"text": "unit testing"}`)
		memory.Add("3", `{"text": "cb insights"}`)

		expectedLen := 3
		expectedCapacity := 10
		s.Equal(memory.Capacity(), expectedCapacity)
		s.Equal(memory.Len(), expectedLen)
	})

	s.Run("MemoryAddExceedsCapacity", func() {
		c := &conf.Config{Slots: 2}
		memory := NewMemoryStorage(c)
		memory.Add("1", `{"text": "ok"}`)
		memory.Add("2", `{"text": "unit testing"}`)

		expectedLen := 2
		expectedCapacity := 2
		s.Equal(memory.Capacity(), expectedCapacity)
		s.Equal(memory.Len(), expectedLen)

		isAdded := memory.Add("3", `{"text": "cb insights"}`)
		s.Equal(false, isAdded)
	})
}

func (s *memoryStorageSuite) TestMemory_Get() {
	s.Run("MemoryGetExistingKey", func() {
		c := &conf.Config{Slots: 10}
		memory := NewMemoryStorage(c)
		firstObject := `{"text": "ok"}`
		memory.Add("1", firstObject)
		memory.Add("2", `{"text": "unit testing"}`)

		object, exist := memory.Get("1")

		s.Equal(exist, true)
		s.Equal(firstObject, object)
	})

	s.Run("MemoryGetNonExistingKey", func() {
		c := &conf.Config{Slots: 10}
		memory := NewMemoryStorage(c)
		firstObject := `{"text": "ok"}`
		memory.Add("1", firstObject)
		memory.Add("2", `{"text": "unit testing"}`)

		object, exist := memory.Get("100")

		s.Equal(exist, false)
		s.Nil(object)
	})
}

func (s *memoryStorageSuite) TestMemory_Delete() {
	s.Run("MemoryDeleteExistingKey", func() {
		c := &conf.Config{Slots: 10}
		memory := NewMemoryStorage(c)
		firstObject := `{"text": "ok"}`
		memory.Add("1", firstObject)
		memory.Add("2", `{"text": "unit testing"}`)

		deletedObject, isDeleted := memory.Delete("1")
		expectedDeleted := true
		s.Equal(expectedDeleted, isDeleted)
		s.Equal(firstObject, deletedObject)
	})

	s.Run("MemoryDeleteNonExistingKey", func() {
		c := &conf.Config{Slots: 10}
		memory := NewMemoryStorage(c)
		memory.Add("1", `{"text": "ok"}`)
		memory.Add("2", `{"text": "unit testing"}`)

		deletedObject, isDeleted := memory.Delete("100")
		expectedDeleted := false
		s.Equal(expectedDeleted, isDeleted)
		s.Nil(deletedObject)
	})
}

func (s *memoryStorageSuite) TestMemory_Front() {
	s.Run("NilOnEmptyStorage", func() {
		c := &conf.Config{Slots: 10}
		m := NewMemoryStorage(c)
		s.Nil(m.Front())
	})

	s.Run("NonEmptyStorage", func() {
		c := &conf.Config{Slots: 10}
		m := NewMemoryStorage(c)
		m.Add("1", `{"text": "front"}`)
		s.NotNil(m.Front())
	})
}

func (s *memoryStorageSuite) TestMemory_Back() {
	s.Run("NilOnEmptyStorage", func() {
		c := &conf.Config{Slots: 10}
		m := NewMemoryStorage(c)
		s.Nil(m.Back())
	})

	s.Run("NonEmptyStorage", func() {
		c := &conf.Config{Slots: 10}
		m := NewMemoryStorage(c)
		m.Add("1", `{"text": "front"}`)
		s.NotNil(m.Back())
	})
}
