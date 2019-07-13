package model

import "time"

//Account struct
type Account struct {
	Login    string
	Password string
}

//NewAccount initializes a new Account
func NewAccount() *Account {
	return &Account{}
}

//User struct
type User struct {
	Accounts  []Account
	Email     string
	Roles     []string
	Active    bool
	LastLogin time.Time
}

//NewUser initializes a new User
func NewUser() *User {
	return &User{}
}
