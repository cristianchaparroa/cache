package datasources

import (
	"github.com/stretchr/testify/suite"
	"os"
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

func (s *memoryStorageSuite) TestMemory_new_with_bad_slot_variable() {
	os.Setenv("SLOTS", "it-should-be-number")
	defer os.Clearenv()

	defer func() {
		if r := recover(); r == nil {
			s.Fail("Call the constructor dont panic")
		}
	}()

	_ = NewMemoryStorage()
}

func (s *memoryStorageSuite) TestMemory_init_capacity() {
	os.Setenv("SLOTS", "10")
	defer os.Clearenv()
	expectedLen := 0
	expectedCapacity := 10

	memory := NewMemoryStorage()
	s.Equal(memory.Len(), expectedLen)
	s.Equal(memory.Capacity(), expectedCapacity)
}

func (s *memoryStorageSuite) TestMemory_capacity_by_default() {
	memory := NewMemoryStorage()
	expectedLen := 0
	s.Equal(memory.Len(), expectedLen)
	s.Equal(memory.Capacity(), defaultCapacity)
}

func (s *memoryStorageSuite) TestMemory_Add_repeated_keys() {
	os.Setenv("SLOTS", "10")
	defer os.Clearenv()

	memory := NewMemoryStorage()
	memory.Add("1", `{"text": "ok"}`)
	memory.Add("1", `{"text": "unit testing"}`)
	memory.Add("1", `{"text": "cb insights"}`)

	expectedLen := 1
	expectedCapacity := 10

	s.Equal(memory.Capacity(), expectedCapacity)
	s.Equal(memory.Len(), expectedLen)
}

func (s *memoryStorageSuite) TestMemory_add_simple() {
	os.Setenv("SLOTS", "10")
	defer os.Clearenv()

	memory := NewMemoryStorage()
	memory.Add("1", `{"text": "ok"}`)
	memory.Add("2", `{"text": "unit testing"}`)
	memory.Add("3", `{"text": "cb insights"}`)

	expectedLen := 3
	expectedCapacity := 10
	s.Equal(memory.Capacity(), expectedCapacity)
	s.Equal(memory.Len(), expectedLen)
}

func (s *memoryStorageSuite) TestMemory_Add_exceeds_capacity() {
	os.Setenv("SLOTS", "2")
	defer os.Clearenv()

	memory := NewMemoryStorage()
	memory.Add("1", `{"text": "ok"}`)
	memory.Add("2", `{"text": "unit testing"}`)

	expectedLen := 2
	expectedCapacity := 2
	s.Equal(memory.Capacity(), expectedCapacity)
	s.Equal(memory.Len(), expectedLen)

	isAdded := memory.Add("3", `{"text": "cb insights"}`)
	s.Equal(false, isAdded)
}

func (s *memoryStorageSuite) TestMemory_Get_existing_key() {
	memory := NewMemoryStorage()
	firstObject := `{"text": "ok"}`
	memory.Add("1", firstObject)
	memory.Add("2", `{"text": "unit testing"}`)

	object, exist := memory.Get("1")

	s.Equal(exist, true)
	s.Equal(firstObject, object)
}

func (s *memoryStorageSuite) TestMemory_Get_non_existing_key() {
	memory := NewMemoryStorage()
	firstObject := `{"text": "ok"}`
	memory.Add("1", firstObject)
	memory.Add("2", `{"text": "unit testing"}`)

	object, exist := memory.Get("100")

	s.Equal(exist, false)
	s.Nil(object)
}

func (s *memoryStorageSuite) TestMemory_Delete_existing_key() {
	memory := NewMemoryStorage()
	firstObject := `{"text": "ok"}`
	memory.Add("1", firstObject)
	memory.Add("2", `{"text": "unit testing"}`)

	deletedObject, isDeleted := memory.Delete("1")
	expectedDeleted := true
	s.Equal(expectedDeleted, isDeleted)
	s.Equal(firstObject, deletedObject)
}

func (s *memoryStorageSuite) TestMemory_Delete_non_existing_key() {
	memory := NewMemoryStorage()
	memory.Add("1", `{"text": "ok"}`)
	memory.Add("2", `{"text": "unit testing"}`)

	deletedObject, isDeleted := memory.Delete("100")
	expectedDeleted := false
	s.Equal(expectedDeleted, isDeleted)
	s.Nil(deletedObject)
}

func (s *memoryStorageSuite) TestMemory_Front() {
	s.Run("NilOnEmptyStorage", func() {
		m := NewMemoryStorage()
		s.Nil(m.Front())
	})

	s.Run("NilOnEmptyStorage", func() {
		m := NewMemoryStorage()
		m.Add("1", `{"text": "front"}`)
		s.NotNil(m.Front())
	})
}

func (s *memoryStorageSuite) TestMemory_Back() {
	s.Run("NilOnEmptyStorage", func() {
		m := NewMemoryStorage()
		s.Nil(m.Back())
	})

	s.Run("NilOnEmptyStorage", func() {
		m := NewMemoryStorage()
		m.Add("1", `{"text": "front"}`)
		s.NotNil(m.Back())
	})
}
