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
)

type User struct {
	id        string
	username  string
	email     string
	name      string
	status    bool
	role      UserRole
	password  string
	createdAt time.Time
	updatedAt time.Time
}

type NewUser struct {
	username string
	email    string
	name     string
	status   bool
	role     UserRole
	password string
}

func (newUser *NewUser) toDomainMapper() *User {
	return &User{
		username: newUser.username,
		name:     newUser.name,
		email:    newUser.email,
		role:     newUser.role,
	}
}

type Service interface {
	GetAll() (*[]User, error)
	GetById(id string) (*User, error)
	Create(newUser *NewUser) (*User, error)
	GetOneByMap(userMap map[string]interface{}) (*User, error)
	Delete(id string) error
	Update(id string, userMap map[string]interface{}) (*User, error)
}

func (r UserRole) isValidRole() error {
	switch r {
	case ADMIN, CLIENT, DELIVER:
		return nil
	}

	return errors.New("Invalod Userrole")
}
