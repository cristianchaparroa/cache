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

	err := memory.Add("3", `{"text": "cb insights"}`)
	if err == nil {
		s.Fail("expected an error but get nil")
	}
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
