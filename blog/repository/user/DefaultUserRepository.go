package blog

import "fmt"

type userRepository struct {
}

func newUserRepository() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) FindAll() []*UserEntity {
	return []*UserEntity{
		{
			username: "user123",
			password: "pass123",
		},
	}
}

func (repository *userRepository) FindUserWithUsernameAndPassword(username string, password string) (*UserEntity, error) {
	for _, user := range repository.FindAll() {
		if user.username == username && user.password == password {
			return &UserEntity{
				username: username,
				password: password,
			}, nil
		}
	}
	return nil, fmt.Errorf("No user found with this username and password.")
}
