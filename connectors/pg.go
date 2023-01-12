package connectors

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	errw := godotenv.Load(".env")
	if errw != nil {

	}

	var err error
	dsn := os.Getenv("MYSQL_DSN_CONNECTION_DETAILS")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("disini")
		panic(err)
	}

	log.Println("Connected to Databases")

}
