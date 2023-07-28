package services

import (
	model "example/socket/Model"
)

func ingectDb() *map[string]model.User {
	db := make(map[string]model.User )
	return &db
}