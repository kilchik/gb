package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"l3/cmd/mockery/mocks"
	"testing"
)




//mockery --name DB

func TestCalculator_ProcessOrder(t *testing.T) {
	dbMock := &mocks.DB{}
	calc := &Calculator{db: dbMock}

	// Проверяем  аргументы и заодно явно специфицируем возвращаемые значения
	dbMock.On("GetUserName", uint64(42)).Return("Bob", nil)
	//dbMock.On("GetOrderItems", uint64(100500)).Return([]uint64{100, 200, 250}, nil)

	//dbMock.On("GetUserName", mock.Anything).Return("Bob", nil)
	//dbMock.On("GetOrderItems", mock.Anything).Return([]uint64{100, 200, 250}, nil)

	res, err := calc.ProcessOrder(42, 100500)
	require.NoError(t, err)
	assert.Equal(t, "user Bob spent $550", res)
}
