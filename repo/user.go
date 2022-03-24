package repo

import (
	"context"
	"github.com/suaas21/pathao/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*model.User, error)

	QueryUsers(ctx context.Context, conditions map[string]interface{}) ([]*model.User, error)
}
