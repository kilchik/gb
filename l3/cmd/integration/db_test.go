// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetUsers(t *testing.T)  {
	dsn := os.Getenv("DB_TEST_DSN")
	require.NotEmpty(t, dsn)

	db := &DBImpl{}
	insertUser(t, 42, "Bob")
	assert.Equal(t, []int64{42}, db.GetUsers())
}

func insertUser(t *testing.T, id int64, name string) {
	t.Helper()
	// sql insert query
}