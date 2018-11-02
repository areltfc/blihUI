// Go project by arthur
// blihUI
// 2018

package user

import (
	"blihUI/pkg/token"
	"github.com/mewbak/gopass"
)

type User struct {
	email, token string
}

func New(email string) *User {
	password, err := gopass.GetPass("Mot de passe bocal : ")
	if err != nil {
		panic(err)
	}
	t := token.Token(password)
	return &User{email: email, token: t.ToSha512()}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Token() string {
	return u.token
}
