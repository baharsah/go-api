package migration

import (
	"go-api/connectors"
	"go-api/models"

	log "github.com/sirupsen/logrus"
)

func RunMigration() {

	err := connectors.DB.AutoMigrate(&models.Tickets{}, models.Transaction{}, models.User{}, models.Stations{})

	if err != nil {
		log.Println(err)
		panic("Migras Gagal")
	} else {
		log.Println("Migrasi Berhasil")
	}

}
