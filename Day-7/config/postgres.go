package config

import (
	"clean_arch/entity"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("db_url")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected")
}
func Migrate() {
	DB.AutoMigrate(&entity.User{})
}
