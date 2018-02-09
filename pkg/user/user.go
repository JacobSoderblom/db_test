package user

import "errors"

var (
	ErrUnknown = errors.New("unknown user")
)

type UserID int

type User struct {
	ID UserID
	Name string
	Email string
	Password string
	Address Address
}

func (u *User) SetName(name string) {
	u.Name = name
}
func (u *User) SetEmail(email string) {
	u.Email = email
}
func (u *User) SetAddress(addr Address) {
	u.Address = addr
}

type Address struct {
	Street string
	City string
}

type Repository interface {
	Store(*User) error
	Find(UserID) (*User, error)
	FindList() (*[]byte, error)
}

func NewUser(name, email, password string, addr Address) *User {
	return &User{
		Name: name,
		Email: email,
		Password: password,
		Address: addr,
	}
}