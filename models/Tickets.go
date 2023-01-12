package models

import (
	"time"

	"gorm.io/gorm"
)

type Tickets struct {
	gorm.Model
	TrainName    string
	StartDate    time.Time
	StationStart Stations
	StationEnd   Stations
	Price        int
	Qty          int
}
