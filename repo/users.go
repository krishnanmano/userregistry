package repo

import (
	"golang.org/x/net/context"
	"userregistry/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . UsersRepo
type UsersRepo interface {
	Save(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User, userName string) error
	GetByUserName(ctx context.Context, userName string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Close() error
}
