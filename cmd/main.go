package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jakefish18/GolangHttpServer/cmd/api"
	"github.com/jakefish18/GolangHttpServer/config"
	"github.com/jakefish18/GolangHttpServer/db"
	"log"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start")
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Sucessfully connected")
}
