package service

import (
	"strconv"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pakawatkung/gogorm/repository"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userSrv repository.UserRepository
}

func NewUserService(userSrv repository.UserRepository) UserService  {
	return userService{userSrv: userSrv}
}

func (s userService) CreateTableUser() error {

	err := s.userSrv.CreateUserTable()
	if err != nil {
		return err
	}

	return nil

}

func (s userService) SignUpUser(userReq UserRequire) (*UserResponse, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		return nil, err
	}

	userRepo := repository.CreateUser{
		Username: userReq.Username,
		Password: string(password),
	}

	userRes, err := s.userSrv.InsertUserData(userRepo)
	if err != nil {
		return nil, err
	}

	userResponse := UserResponse{
		ID: userRes.ID,
		Username: userRes.Username,
		Password: userRes.Password,
	}

	return &userResponse, nil

}

func (s userService) SignInUser(userReq UserRequire) (*TokenResponse, error) {

	idUser, err := s.userSrv.GetOneUser(userReq.Username)
	if err != nil {
		return nil, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(idUser.Password), []byte(userReq.Password))
	if err != nil {
		return nil, err
	}

	claims := jwt.StandardClaims{
		Issuer: strconv.Itoa(idUser.ID),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(viper.GetString("key.jwt")))
	if err != nil {
		return nil, err
	}

	tokenResponse := TokenResponse{
		Username: idUser.Username,
		Password: idUser.Password,
		Token: token,
	}

	return &tokenResponse, nil
}

func (s userService) GetAllUser() ([]UserResponse, error) {

	datas, err := s.userSrv.GetAllUser()
	if err != nil {
		return nil, err
	}

	userResponse := []UserResponse{}

	for _, data := range datas {
		user := UserResponse{
			ID: data.ID,
			Username: data.Username,
			Password: data.Password,
		}
		userResponse = append(userResponse, user)
	}

	return userResponse, nil
}

func (s userService) GetOneUser(uname string) (*UserResponse, error) {

	user, err := s.userSrv.GetOneUser(uname)
	if err != nil {
		return nil, nil
	}

	userResponse := UserResponse{
		ID: user.ID,
		Username: user.Username,
		Password: user.Password,
	}

	return &userResponse, nil
}