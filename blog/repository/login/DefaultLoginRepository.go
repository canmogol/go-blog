package blog

type defaultLoginRepository struct {
}

func newDefaultLoginRepository() LoginRepository {
	return &defaultLoginRepository{}
}

func (repository *defaultLoginRepository) Login(username string, password string) bool {
	return true
}
