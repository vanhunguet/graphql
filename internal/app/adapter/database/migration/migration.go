package migration

import (
	"github.com/vanhunguet/GrabhQL/graph/model"
	"gorm.io/gorm"
	"log"
)

func Migrations(db *gorm.DB) {
	err := db.Debug().AutoMigrate(
		&model.User{},
		&model.Video{},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Migration successfully.")
}
