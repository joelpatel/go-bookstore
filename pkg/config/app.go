package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// NOTE: whole point of this package is to return this DB to the overall package/repo

func Connect() {
	dsn := "host=localhost dbname=gobookstore port=5400"
	p_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = p_db
}

func GetDB() *gorm.DB {
	return db
}
