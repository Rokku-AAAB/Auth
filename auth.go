package auth

import (
	"github.com/Tnze/go-mc/yggdrasil"
)

type AuthResult struct {
	UUID string
	Name string
}

func VerifyAuth(user, password string) (AuthResult, error) {
	resp, err := yggdrasil.Authenticate(user, password)

	if err != nil {
		return AuthResult{}, err
	}

	id, name := resp.SelectedProfile()

	return AuthResult{UUID: id, Name: name}, nil
}
