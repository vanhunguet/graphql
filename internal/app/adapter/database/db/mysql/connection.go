package mysql

import (
	"fmt"
	"github.com/vanhunguet/GrabhQL/internal/app/adapter/database/migration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "vanhung12345678#", "127.0.0.1", "3306", "grabhql")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database mysql")
		return nil
	}
	migration.Migrations(db)
	DB = db
	return DB
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection mysql")
	}
	err = dbSQL.Close()
	if err != nil {
		panic("Failed to close database database mysql")
	}
}
