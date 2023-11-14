package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/murilo-toddy/pixel/domain/model"
    _ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(file)

	err := godotenv.Load(basepath + "/../../.env")
	if err != nil {
		log.Fatalf("Error loading .env")
	}
}

func Connect(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv("dsn")
		db, err = gorm.Open(os.Getenv("dbType"), dsn)
	} else {
		dsn = os.Getenv("dsnTest")
		db, err = gorm.Open(os.Getenv("dbType"), dsn)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
