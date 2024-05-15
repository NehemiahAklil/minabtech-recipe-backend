package entity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

func (u *User) SetPassword(password string) error {
	hashedPasssword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPasssword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	if u.Password != "" {
		fmt.Println("Hashed Password : " + u.Password)
	} else {
		fmt.Println("Password is empty")
	}
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
