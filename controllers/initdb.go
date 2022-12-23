package controllers

import (
	"github.com/joho/godotenv"
	"go-techradar/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("DATABASE")
	if val == "POSTGRES" {
		log.Println("using postgres")
		dsn := "host=localhost user=zacktest password=zacktest dbname=gotodo port=5434 sslmode=disable TimeZone=Asia/Kuala_Lumpur"
		dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		log.Println("Connection to DB success !!")

		db = dbConnect
	} else if val == "SQLITE" {
		log.Println("using sqlite")

		dbConnect, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = dbConnect
	}
	// Migrate the schema
	if err = db.AutoMigrate(&model.Blip{}); err != nil {
		panic("failed to create schema")
	}

	log.Println("Auto Migrate DB success !!")
}
