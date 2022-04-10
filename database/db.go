package database

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
	DB *gorm.DB
	err error
)

func Conn() {
	string := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(string))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados", err)
	}
}