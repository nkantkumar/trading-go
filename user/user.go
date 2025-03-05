package user

import (
	"fmt"
)

type User struct {
	Name  string
	email string
}

// GetEmail returns the user's email
func (u *User) GetEmail() string {
	return u.email
}

// SetEmail sets the user's email with basic validation
func (u *User) SetEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	u.email = email
	return nil
}
