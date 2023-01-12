package repositories

import (
	"go-api/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepo interface {
	CreateUser(models.User) (models.User, error)
	SelectUser(models.User) (models.User, error)
}

func RepoAuth(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) CreateUser(user models.User) (models.User, error) {
	e := r.db.Create(&user).Error

	return user, e
}

func (r *repo) SelectUser(data models.User) (models.User, error) {

	var user models.User
	err := r.db.Debug().Where("username = ? ", data.Username).First(&user).Error
	log.Println(user)
	return user, err

}
