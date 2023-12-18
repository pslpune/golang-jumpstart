package auth

import "fmt"

type UserType uint8

const (
	Developer UserType = iota
	Admin
	Manager
	Guest
)

type User interface {
	Details() string
}

type Auth interface {
	Exists() (bool, error)
	Login(p string) (bool, error)
}

// AnyUser : just about any user that can be authenticated and dispatched over http in json format.
type AnyUser struct {
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Loc    string   `json:"location"`
	Phone  string   `json:"phone"`
	Passwd string   `json:"passwd"`
	UType  UserType `json:"utype"`
}

func (au *AnyUser) Details() string {
	return fmt.Sprintf("%s-%s", au.Email, au.Name)
}
func (au *AnyUser) Exists() (bool, error) {
	return true, nil
}
func (au *AnyUser) Login(p string) (bool, error) {
	return true, nil
}

func NewUser(name, email, loc, phone string) User {
	return &AnyUser{
		Name:  name,
		Email: email,
		Loc:   loc,
		Phone: phone,
		UType: Guest,
	}
}
