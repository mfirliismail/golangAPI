package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConfig() *gorm.DB {

	dsn := "host=localhost user=postgres password=postgres dbname=golang_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed connect to db")
	}

	// db.autoMigrate()

	return db
}

func CloseDatabaseConfig(db *gorm.DB) {
	dbPostgre, err := db.DB()

	if err != nil {
		panic("failed to close db")
	}

	dbPostgre.Close()

}
