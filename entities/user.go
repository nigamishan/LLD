package entities

import "SnakesAndLadder/common"

type User struct {
	Id              int
	Name            string
	State           common.PlayerState
	CurrentPosition int
}

func CreateUser(id int, name string) User {
	return User{
		Id:              id,
		Name:            name,
		State:           common.ACTIVE,
		CurrentPosition: 1,
	}
}
