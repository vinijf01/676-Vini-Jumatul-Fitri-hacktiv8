package databases

import (
	"ass2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	//todo
	// dsn
	dsn := "root:@tcp(127.0.0.1:3306)/db_assignment2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//open database
	db = db.Debug()
	autoMigrate()
}

func GetDb() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(&models.Order{}, &models.Items{})
}
