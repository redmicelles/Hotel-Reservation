package db

import "github.com/redmicelles/Hotel-Reservation/types"

type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStore struct{}
