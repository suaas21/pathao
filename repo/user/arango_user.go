package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/suaas21/pathao/infra"
	"github.com/suaas21/pathao/logger"
	"github.com/suaas21/pathao/model"
	"github.com/suaas21/pathao/repo"
)

const CollectionUser = "User"

type userRepo struct {
	db  infra.ArangoDB
	log logger.StructLogger
}

func NewArangoUserRepository(db infra.ArangoDB, lgr logger.StructLogger) repo.UserRepository {
	return &userRepo{
		db:  db,
		log: lgr,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user model.User) error {
	if err := u.db.CreateDocument(ctx, CollectionUser, &user); err != nil {
		return err
	}

	return nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user model.User) error {
	if err := u.db.UpdateDocument(ctx, CollectionUser, user.ID, &user); err != nil {
		return err
	}

	return nil
}

func (u *userRepo) DeleteUser(ctx context.Context, id string) error {
	return u.db.RemoveDocument(ctx, CollectionUser, id)

}

func (u *userRepo) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := u.db.ReadDocument(ctx, CollectionUser, id, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) QueryUsers(ctx context.Context, conditions map[string]interface{}) ([]*model.User, error) {

	// TODO: if conditions provided then will be added to filter in query
	userQuery := fmt.Sprintf(`FOR x IN User RETURN x`)

	res, err := u.db.Query(ctx, userQuery, nil)
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, 0)
	dataBytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
