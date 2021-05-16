package blog

type LoginService interface {
	Login(username string, password string) (bool, error)
}
