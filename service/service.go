package service

import (
	"context"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/model"
)

type Controller interface {
	Insert(ctx context.Context, u *model.User) error
	User(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, u model.User) (model.User, error)
	Delete(ctx context.Context, id int) error
	All(ctx context.Context) ([]model.User, error)
}

type Service struct {
	DB Controller
}

func New(controller Controller) Service {
	return Service{DB: controller}
}
