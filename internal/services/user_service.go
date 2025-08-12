package services

import (
	"github.com/tharunn0/gin-server-gorm/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryIf interface {
	RegisterUser(user *models.User) error
}

type UserService struct {
	repository UserRepositoryIf
}

func NewUserService(repo UserRepositoryIf) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (u *UserService) Create(rq *models.RegisterReq) error {
	var user *models.User

	hashedpass, er := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if er != nil {
		return er
	}

	user = &models.User{
		Email:      rq.Email,
		Username:   rq.Username,
		Firstname:  rq.Firstname,
		Lastname:   rq.Lastname,
		IsAdmin:    false,
		IsVerified: false,
		IsBlocked:  false,
		Password:   string(hashedpass),
	}

	er = u.repository.RegisterUser(user)
	if er != nil {
		return er
	}
	return nil
}
