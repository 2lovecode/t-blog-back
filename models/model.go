package models

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
}

var TankDb *gorm.DB

func InitTankDb(db *gorm.DB) {
	TankDb = db

	TankDb.AutoMigrate(&Module{})
	TankDb.AutoMigrate(&Tag{})
}
