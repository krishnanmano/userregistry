package repo

import (
	"golang.org/x/net/context"
	"userregistry/models"
)

type UsersRepo interface {
	Save(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User, userName string) error
	GetByUserName(ctx context.Context, userName string) (*models.User, error)
	Close() error
}
