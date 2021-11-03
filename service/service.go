package service

import (
	"context"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/model"
)

// Controller is the manipulation tool of the db from web api side.
type Controller interface {
	Insert(ctx context.Context, u *model.User) error
	User(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, u model.User) error
	Delete(ctx context.Context, id int) error
	All(ctx context.Context) ([]model.User, error)
}

// Conn is the connection to Controller interface.
type Conn struct {
	DB Controller
}
