package main

import (
	"database/sql"
	errors2 "errors"
	"github.com/pkg/errors"
	"log"
)

type UserDao struct {

}

type User struct {
	Name string
	Age int
}

func (user *UserDao) findOne(id string) (*User, error)  {
	err := sql.ErrNoRows
	if errors2.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrapf(err, "not found by id : %s", id)
	}

	return &User{"A", 1}, nil
}

func (user *UserDao) findAll() ([]*User, error) {
	err := sql.ErrNoRows

	if errors2.Is(err, sql.ErrNoRows) {
		log.Println("query findAll result is Empty")
		return []*User{}, nil
	}

	return []*User{}, nil
}
