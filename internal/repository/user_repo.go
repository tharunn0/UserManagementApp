package repository

import (
	"fmt"

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
		return fmt.Errorf("repository.RegisterUser : %w/n", er)
	}
	return nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	tx := repo.pgdb.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, fmt.Errorf("repository.LoginUser : %w/n", tx.Error)
	}
	return user, nil
}

func (repo *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	tx := repo.pgdb.Where("is_admin = ?", false).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
