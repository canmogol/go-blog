package blog

import (
	repositoryUser "github.com/canmogol/godi/blog/repository/user"
)

type userLoginRepository struct {
	userRepository repositoryUser.UserRepository
}

func newUserLoginRepository(userRepository repositoryUser.UserRepository) LoginRepository {
	return &userLoginRepository{
		userRepository: userRepository,
	}
}

func (repository *userLoginRepository) Login(username string, password string) bool {
	_, error := repository.userRepository.FindUserWithUsernameAndPassword(username, password)
	if error != nil {
		return false
	}
	return true
}
