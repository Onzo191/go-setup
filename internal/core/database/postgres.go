package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(host, port, user, password, dbname string) (*gorm.DB, error) {
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
