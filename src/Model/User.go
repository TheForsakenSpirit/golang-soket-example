package model

import (
	"math/rand"
)

type User struct {
	Key      string `json:"key"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"-"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateKey(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}
