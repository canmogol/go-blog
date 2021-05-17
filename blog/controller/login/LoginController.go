package blog

type LoginController interface {
	Login(loginRequest *LoginRequest) *LoginResponse
}
