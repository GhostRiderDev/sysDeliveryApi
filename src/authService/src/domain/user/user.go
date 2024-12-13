package user

import (
	"errors"
	"time"
)

type UserRole string

const (
	ADMIN   UserRole = "ADMIN"
	CLIENT  UserRole = "CLIENT"
	DELIVER UserRole = "DELIVER"
	alguma = "hello2"
)

type User struct {
	id       string
	username string
	email    string
	name     string
	status   bool
	role     UserRole
	password string
	createdAt time.Time
	updatedAt time.Time
}

func (r UserRole) isValidRole() error {
	switch r {
	case ADMIN, CLIENT, DELIVER:
		return nil
	}

	return errors.New("Invalod Userrole")
}
