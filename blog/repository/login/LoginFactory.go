package blog

import (
	"fmt"

	userRepository "github.com/canmogol/godi/blog/repository/user"
)

func CreateLoginRepository(repositoryType string, userRepository userRepository.UserRepository) (LoginRepository, error) {
	if repositoryType == "default" {
		return newDefaultLoginRepository(), nil
	}
	if repositoryType == "user" {
		return newUserLoginRepository(userRepository), nil
	}
	return nil, fmt.Errorf("Could not create LoginRepository, unknown repository type passed: '%s'", repositoryType)
}
