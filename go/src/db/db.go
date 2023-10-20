package db

import (
	"log"
	"os"

	"github.com/cezar-tech/fullcycle01grpcgorm/go/src/repository"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env files %v", err)
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	dsn = os.Getenv("dsn")
	db, err = gorm.Open(os.Getenv("dbType"), dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&repository.ProductDAO{})
	}

	return db
}
