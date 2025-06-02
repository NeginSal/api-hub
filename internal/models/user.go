package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`  
	PasswordHash string `json:"-"`
}

var DummyUser = func() User {
	password := "1234"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return User{
		Username:     "admin",
		PasswordHash: string(hash),
	}
}()
