package database

import "github.com/google/uuid"

type ID = uuid.UUID

type UserSchema struct {
	Id        ID     `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type UserDB map[ID]UserSchema

type DB struct {
	UsersDB UserDB
}

func NewDB() *DB {
	return &DB{
		UsersDB: UserDB{},
	}
}
