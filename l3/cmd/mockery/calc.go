package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type DB interface {
	GetUserName(uid uint64) (string, error)
	GetOrderItems(oid uint64) ([]uint64, error)
}

type Calculator struct {
	db DB
}

func (c *Calculator) ProcessOrder(userId, orderId uint64) (string, error) {
	name, err := c.db.GetUserName(userId)
	if err != nil {
		return "", errors.Wrap(err, "get user name from db")
	}

	bought, err := c.db.GetOrderItems(orderId)
	if err != nil {
		return "", errors.Wrap(err, "get order items from db")
	}

	return fmt.Sprintf("user %s spent $%d", name, sum(bought)), nil
}

