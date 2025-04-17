package database

import (
	"database/sql"
	"fmt"
	
	_ "github.com/lib/pq"

)

type Database struct {
	Db *sql.DB
}

func NewDataBase() *Database {
	connSrt := "postgres://postgres:294332@localhost:5432/test?sslmode=disable"
	db, err := sql.Open("postgres", connSrt)
	if err != nil {
		fmt.Printf("Error al conectar con la base de datos: %v\n", err)
	}
	return &Database{
		Db: db,
	}
}
