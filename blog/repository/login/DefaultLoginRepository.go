package blog

type defaultLoginRepository struct {
}

func (repository *defaultLoginRepository) Login(username string, password string) bool {
	return true
}

func newDefaultLoginRepository() LoginRepository {
	return &defaultLoginRepository{}
}
