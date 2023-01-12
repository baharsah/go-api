package repositories

import (
	"errors"
	"go-api/models"

	"gorm.io/gorm"
)

type StationRepo interface {
	CreateStation(models.Stations) (models.Stations, error)
	GetStation() ([]models.Stations, error)
}

func RepoStation(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) CreateStation(station models.Stations) (models.Stations, error) {
	var datastations models.Stations
	var stationCount int64

	r.db.Where("station = ?", station.Station).First(&datastations).Count(&stationCount)
	if stationCount > 0 {
		return datastations, errors.New("stasiun Sudah ada")
	}
	e := r.db.Create(&station).Error
	return station, e
}

func (r *repo) GetStation() ([]models.Stations, error) {

	var stations []models.Stations
	e := r.db.Find(&stations).Error

	return stations, e

}
