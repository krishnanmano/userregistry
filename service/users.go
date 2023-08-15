package service

import (
	"context"
	"fmt"
	"time"
	"userregistry/models"
	"userregistry/repo"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User, userName string) error
	Get(ctx context.Context, userName string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
}

type userService struct {
	usersRepo repo.UsersRepo
}

func NewUserService(usersRepo repo.UsersRepo) UserService {
	return &userService{usersRepo: usersRepo}
}

func (u *userService) Create(ctx context.Context, user *models.User) error {
	expiryTime := time.Unix(user.ExpiryDate, 0)
	fmt.Println("Expiry Date", expiryTime)
	return u.usersRepo.Save(ctx, user)
}

func (u *userService) Update(ctx context.Context, user *models.User, userName string) error {
	expiryTime := time.Unix(user.ExpiryDate, 0)
	fmt.Println("Expiry Date", expiryTime)
	return u.usersRepo.Update(ctx, user, userName)
}

func (u *userService) Get(ctx context.Context, userName string) (*models.User, error) {
	return u.usersRepo.GetByUserName(ctx, userName)
}

func (u *userService) GetAll(ctx context.Context) ([]models.User, error) {
	return u.usersRepo.GetAll(ctx)
}
