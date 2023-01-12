package repositories

import (
	"go-api/models"

	"gorm.io/gorm"
)

type TicketRepo interface {
	CreateTicket(models.Tickets) (models.Tickets, error)
	GetTickets() ([]models.Tickets, error)
}

func RepoTicket(db *gorm.DB) *repo {

	return &repo{db}
}

func (r *repo) CreateTicket(tickets models.Tickets) (models.Tickets, error) {

	e := r.db.Create(&tickets).Error
	return tickets, e

}

func (r *repo) GetTickets() ([]models.Tickets, error) {
	var tickets []models.Tickets
	e := r.db.Find(&tickets).Error
	return tickets, e
}
