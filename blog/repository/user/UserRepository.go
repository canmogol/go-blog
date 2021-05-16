package blog

type UserRepository interface {
	FindAll() []*UserEntity
	FindUserWithUsernameAndPassword(username string, password string) (*UserEntity, error)
}
