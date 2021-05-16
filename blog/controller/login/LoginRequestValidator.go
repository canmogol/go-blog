package blog

import (
	"fmt"
	"strings"
)

type loginRequestValidator struct {
}

func (validator *loginRequestValidator) validate(loginRequest *LoginRequest) (bool, error) {
	if strings.TrimSpace(loginRequest.username) == "" {
		return false, fmt.Errorf("Username cannot be empty.")
	}
	if strings.TrimSpace(loginRequest.password) == "" {
		return false, fmt.Errorf("Password cannot be empty.")
	}
	return true, nil
}
