package repository

import (
	"github.com/tharunn0/gin-server-gorm/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	pgdb *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		pgdb: db,
	}
}

func (repo *UserRepository) RegisterUser(user *models.User) error {
	er := repo.pgdb.Create(user).Error
	if er != nil {
		return er
	}
	return nil
}
