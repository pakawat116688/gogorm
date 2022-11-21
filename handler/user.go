package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/gogorm/service"
)

type userApi struct {
	userAPI service.UserService
}

func NewUserApi(userAPI service.UserService) userApi  {
	return userApi{userAPI: userAPI}
}

func (a userApi) CreateTableApi(c *fiber.Ctx) error {

	err := a.userAPI.CreateTableUser()
	if err != nil{
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusCreated).SendString("create user table success...")

}

func (a userApi) SignUpUserApi(c *fiber.Ctx) error {

	userReq := service.UserRequire{}
	err := c.BodyParser(&userReq)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userResponse, err := a.userAPI.SignUpUser(userReq)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(userResponse)

}

func (a userApi) SignInUserApi(c *fiber.Ctx) error  {

	userReq := service.UserRequire{}
	err := c.BodyParser(&userReq)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}
	token, err := a.userAPI.SignInUser(userReq)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "username or password is incorrect !!!")
	}

	return c.Status(fiber.StatusOK).JSON(token)

}

func (a userApi) GetAllUserApi(c *fiber.Ctx) error {

	users, err := a.userAPI.GetAllUser()
	if  err != nil {
		return fiber.NewError(fiber.StatusNotFound, "can't not get data...")
	}

	return c.Status(fiber.StatusOK).JSON(users)

}

func (a userApi) GetOneUserApi(c *fiber.Ctx) error {

	user, err := a.userAPI.GetOneUser(c.Params("username"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "can't not find user...")
	}

	return c.Status(fiber.StatusOK).JSON(user)
	
}