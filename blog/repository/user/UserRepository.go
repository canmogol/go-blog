package blog

type UserRepository interface {
	FindUserWithUsernameAndPassword(username string, password string) (*UserEntity, error)
}
