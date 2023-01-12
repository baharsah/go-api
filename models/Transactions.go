package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Biller   User `gorm:"foreignKey:BillerID"`
	ChildQty int
	AdultQty int
	Ticket   Tickets
}
