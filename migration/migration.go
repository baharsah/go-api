package migration

import (
	"go-api/connectors"
	"go-api/models"
	"log"
)

func RunMigration() {

	err := connectors.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Println(err)
		panic("Migrasi User Gagal")
	} else {
		log.Println("Migrasi User Berhasil")
	}

	err2 := connectors.DB.AutoMigrate(&models.Tickets{})

	if err2 != nil {
		log.Println(err2)
		panic("Migrasi Trip Gagal")
	} else {
		log.Println("Migrasi Trip Berhasil")
	}

	err3 := connectors.DB.AutoMigrate(&models.Transaction{})

	if err3 != nil {
		log.Println(err3)
		panic("Migrasi ImageTrip Gagal")
	} else {
		log.Println("Migrasi ImageTrip Berhasil")
	}
	err4 := connectors.DB.AutoMigrate(&models.Stations{})

	if err4 != nil {
		log.Println(err4)
		panic("Migrasi ImageTrip Gagal")
	} else {
		log.Println("Migrasi ImageTrip Berhasil")
	}

}
