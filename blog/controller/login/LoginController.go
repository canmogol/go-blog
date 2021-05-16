package blog

type LoginController interface {
	Login(loginRequest *LoginRequest) *LoginResponse
	CreateLoginRequest(username string, password string) *LoginRequest
}
