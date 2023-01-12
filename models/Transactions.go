package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	BillerID uint
	Biller   User `gorm:"foreignKey:BillerID"`
	ChildQty int
	AdultQty int
	TicketID uint
	Ticket   Tickets `gorm:"foreignKey:TicketID"`
}

type TransactionRequest struct {
	BillerID uint
	ChildQty int
	AdultQty int
	TicketID uint
}

func (TransactionRequest) TableName() string {
	return "transactions"
}
