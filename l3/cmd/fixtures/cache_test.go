package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CacheLRUTestSuite struct {
	suite.Suite
	cache *CacheLRU
}

func (s *CacheLRUTestSuite) SetupTest() {
	s.cache = NewCacheLRU(5)
	s.cache.Set("foo", 5)
	s.cache.Set("bar", 10)
}

func (s *CacheLRUTestSuite) TestSet() {
	s.T().Run("new key, enough space in cache", func(t *testing.T) {
		s.Require().Nil(s.cache.Set("zoo", 100))
		s.Assert().Equal("zoo", s.cache.MostRecent())
	})

	s.T().Run("new key, not enough space in cache", func(t *testing.T) {
		s.Require().Nil(s.cache.Set("buz", 7))
		s.Assert().Equal("buz", s.cache.MostRecent())
		_, err := s.cache.Get("foo")
		s.Assert().Error(err)
	})
}

func (s *CacheLRUTestSuite) TestGet() {
	s.Require().Equal("foo", s.cache.LeastRecent())
	item, err := s.cache.Get("foo")
	s.Require().NoError(err)
	s.Assert().Equal(5, item)
	s.Assert().Equal("bar", s.cache.LeastRecent())
}

func TestCacheLRUTestSuite(t *testing.T) {
	suite.Run(t, new(CacheLRUTestSuite))
}

