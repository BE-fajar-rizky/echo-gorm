package config

import (
	"fajar/apiMvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connectionString := "root:@tcp(127.0.0.1:3306)/alter_train?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}

// func ini akan dijalankan sebelum func main
// func init() {
// 	fmt.Println("func init berjalan")
// 	initDB()
// 	InitialMigration()
// }
