package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	//creating collection of users
	users  []*User
	nextID = 1
)

//returns user object
func GetUsers() []*User {
	return users
}

//This returns 2 values 1. user object and then error
func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(user User) (User, error) {

	for i, candidate := range users {
		if candidate.ID == user.ID {
			users[i] = &user
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", user.ID)
}

func RemoveUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
