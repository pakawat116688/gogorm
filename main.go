package main

import (
	"os"
	"strings"

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


// func (u User) TableName() string  {
// 	return "User"
// }

func main() {

	os.Remove(viper.GetString("db.file"))
	dialector := sqlite.Open(viper.GetString("db.file"))
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: &sqllogger.SqlLogger{},
		DryRun: true, // true is test
	})
	if err != nil {
		panic(err)
	}


	_ = db
	// db.Migrator().CreateTable(User{})

}


// tx := db.Create(&user) 
// db.Find()

