package managers

import (
	"cache/app/conf"
	"cache/app/datasources"
	"github.com/stretchr/testify/suite"
	"testing"
)

type cacheBuilderSuite struct {
	suite.Suite
}

func TestCacheBuilderSuiteInit(t *testing.T) {
	suite.Run(t, new(cacheBuilderSuite))
}

func (s *cacheBuilderSuite) TestCacheBuilder_Build() {

	c := &conf.Config{
		Slots:  2,
		Policy: "",
	}

	storage := datasources.NewMemoryStorage(c)
	b := NewCacheBuilder(storage)

	s.Run("OlderFistEvictionPolicy", func() {
		cm := b.Build(conf.OlderFistEvictionPolicy)
		s.Equal(conf.OlderFistEvictionPolicy, cm.GetType())
	})

	s.Run("NewestFirstEvictionPolicy", func() {
		cm := b.Build(conf.NewestFirstEvictionPolicy)
		s.Equal(conf.NewestFirstEvictionPolicy, cm.GetType())
	})

	s.Run("RejectEvictionPolicy", func() {
		cm := b.Build(conf.RejectEvictionPolicy)
		s.Equal(conf.RejectEvictionPolicy, cm.GetType())
	})
}
