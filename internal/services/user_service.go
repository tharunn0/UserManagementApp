package services

import (
	"fmt"

	"github.com/tharunn0/gin-server-gorm/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryIf interface {
	RegisterUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUserByUsername(username string) error
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
		return fmt.Errorf("services.GenerateFromPassword : %w/n", er)
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

func (u *UserService) LoginUser(rq *models.LoginReq) (*models.UserProfile, error) {

	var userprofile *models.UserProfile

	email := rq.Email
	password := rq.Password

	user, er := u.repository.GetUserByEmail(email)
	if er != nil {
		return nil, er
	}

	er = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if er != nil {
		return nil, fmt.Errorf("Wrong passowrd !")
	}

	userprofile = &models.UserProfile{
		Username:  user.Username,
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		IsAdmin:   user.IsAdmin,
	}

	return userprofile, nil

}

func (u *UserService) GetAllUsers() ([]*models.UserProfile, error) {
	var userprofiles []*models.UserProfile

	users, er := u.repository.GetAllUsers()
	if er != nil {
		return nil, er
	}

	var up *models.UserProfile
	for _, u := range users {

		up = &models.UserProfile{
			Email:     u.Email,
			Username:  u.Username,
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
		}

		userprofiles = append(userprofiles, up)
	}

	return userprofiles, nil
}

func (s *UserService) GetUserProfile(username string) (*models.UserProfile, error) {

	var userprofile *models.UserProfile

	user, er := s.repository.GetUserByUsername(username)
	if er != nil {
		return nil, er
	}

	userprofile = &models.UserProfile{
		Email:     user.Email,
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}

	return userprofile, nil

}

func (s *UserService) DeleteUserByUsername(username string) error {
	er := s.repository.DeleteUserByUsername(username)
	if er != nil {
		return er
	}
	return nil
}
