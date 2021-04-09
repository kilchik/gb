package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFoo(t *testing.T)  {
	res, err := foo()

	require.NoError(t, err)
	assert.Contains(t, 5, res)
	assert.Contains(t, 3, res)

	require.JSONEq(t, `{"a":42, "b":36}`, `{"b":36, "a":42}`)

	a := 2
	assert.True(t, a > 1 && a < 3)


}