package models

import (
	"time"

	"gorm.io/gorm"
)

type Tickets struct {
	gorm.Model
	TrainName      string
	StartDate      time.Time
	StationStartID uint
	StationStart   Stations `gorm:"foreignKey:StationStartID"`
	StationEndID   uint
	StationEnd     Stations `gorm:"foreignKey:StationEndID"`
	Price          int
	Qty            int
	CreatorID      uint
	Creator        User `gorm:"foreignKey:CreatorID"`
}
