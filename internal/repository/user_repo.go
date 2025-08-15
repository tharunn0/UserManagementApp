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
		return nil, fmt.Errorf("repository.GetUserByEmail : %w/n", tx.Error)
	}
	return user, nil
}
func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User
	tx := repo.pgdb.Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return nil, fmt.Errorf("repository.GetUserByUsername : %w/n", tx.Error)
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

func (repo *UserRepository) DeleteUserByUsername(username string) error {

	tx := repo.pgdb.Table("users").Where("username = ?", username).Delete(nil)

	fmt.Println("the username to delete is ", username)

	if tx.Error != nil || tx.RowsAffected == 0 {
		return fmt.Errorf("repository.DeleteUserByUsername : %w/n", tx.Error)
	}
	fmt.Println("there is no error")
	return nil
}
