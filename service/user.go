package service

type UserResponse struct {
	ID       int    `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequire struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserService interface {
	CreateTableUser() error
	SignUpUser(UserRequire) (*UserResponse, error)
	SignInUser(UserRequire) (*TokenResponse, error)
	GetAllUser() ([]UserResponse, error)
	GetOneUser(string) (*UserResponse, error)
}
