package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/pakawatkung/gogorm/handler"
	"github.com/pakawatkung/gogorm/repository"
	"github.com/pakawatkung/gogorm/service"
	"github.com/pakawatkung/gogorm/sqllogger"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func main() {

	os.Remove(viper.GetString("db.file"))
	dialector := sqlite.Open(viper.GetString("db.file"))
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: &sqllogger.SqlLogger{},
		DryRun: false, // true is test
	})
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserDB(db)
	userService := service.NewUserService(userRepo)
	userApi := handler.NewUserApi(userService)

	err = userService.CreateTableUser()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	app.Post("/signup", userApi.SignUpUserApi)
	app.Post("/signin", userApi.SignInUserApi)
	userGroup := app.Group("/user", jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey: []byte(viper.GetString("key.jwt")),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))
	userGroup.Get("/users", userApi.GetAllUserApi)
	userGroup.Get("/userone/:username", userApi.GetOneUserApi)

	app.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))

}