package models

import "gorm.io/gorm"

type Stations struct {
	gorm.Model
	Station string `gorm:"station_name"`
}
