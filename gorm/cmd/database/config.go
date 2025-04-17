package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func NewDataBase() *Database {
	dsn := "host=localhost user=postgres password=294332 dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("Error abriendo la base de datos: ", err)
	}
	return &Database{
		Db: db,
	}
}
