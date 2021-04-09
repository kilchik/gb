package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// minimock -i DB

func TestCalculator_ProcessOrder(t *testing.T)  {
	dbMock := &DBMock{}
	calc := &Calculator{db: dbMock}

	dbMock.GetUserNameMock.Expect(42).Return("Bob", nil)
	dbMock.GetOrderItemsMock.Expect(100500).Return([]uint64{100, 200, 250}, nil)

	res, err := calc.ProcessOrder(42, 100500)
	require.NoError(t, err)
	assert.Equal(t, "user Bob spent $550", res)
}
