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
