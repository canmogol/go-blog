package blog

type LoginRepository interface {
	Login(username string, password string) bool
}
