package blog

import "fmt"

func CreateUserRepository(repositoryType string) (UserRepository, error) {
	if repositoryType == "default" {
		return newUserRepository(), nil
	}
	return nil, fmt.Errorf("Could not create UserRepository, unknown repository type passed: '%s'", repositoryType)
}
