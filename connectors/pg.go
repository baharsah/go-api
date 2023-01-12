package connectors

import (
	log "github.com/sirupsen/logrus"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	// errw := godotenv.Load(".env")
	// if errw != nil {

	// }

	var err error
	dsn := "host=localhost user=postgres password=backendish dbname=db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("disini")
		panic(err)
	}

	log.Println("Connected to Databases")

}
