package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/murilo-toddy/pixel/domain/model"
	_ "gorm.io/driver/sqlite"
)

func init() {
    _, b, _, _ := runtime.Caller(0)
    basePath := filepath.Dir(b)
    err := godotenv.Load(basePath, "/../../.env")
    if err != nil {
        log.Fatalf("Error loading .env file")
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
        db, err = gorm.Open(os.Getenv("dbTypeTest"), dsn)
    }

    if err != nil {
        log.Fatalf("Error connecting to database %v", err)
        panic(err)
    }

    if os.Getenv("debug") == "true" {
        db.LogMode(true)
    }

    if os.Getenv("autoMigrateDB") == "true" {
        db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
    }

    return db
}

