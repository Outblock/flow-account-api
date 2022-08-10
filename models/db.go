package models

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
)

// Connecting to db
func Connect() *pg.DB {

	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_NAME"),
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	CreateWalletTable(db)
	InitiateDB(db)
	CreateWalletTableMain(db)
	CreateActiveindexTable(db)
	CreateIpLogTable(db)
	return db
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var DbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	DbConnect = db
}
