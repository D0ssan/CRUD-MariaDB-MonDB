package mariadb

import (
	"context"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/model"
	"github.com/pkg/errors"
)

const (
	InsertQuery = `INSERT INTO users (first_name, last_name, specialization, dob) VALUES (?,?,?,?)` // return id
	SelectQuery = `SELECT * FROM users users WHERE id = ?`
	UpdateQuery = `UPDATE users SET first_name=?, last_name=?, specialization=?, dob=? WHERE id= ?`
	DeleteQuery = `DELETE FROM users WHERE id = ?`
)

func (m MariaDB) Insert(ctx context.Context, u *model.User) error {
	query := m.db.Rebind(InsertQuery) //help avoid SQL injection
	res, err := m.db.ExecContext(ctx, query, u.FirstName, u.LastName, u.Specialization, u.DOB)
	if err != nil {
		return errors.Wrap(err, "error executing the insert query into user table")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "error parsing the last id from user table")
	}
	u.ID = id
	return nil
}

func (m MariaDB) User(ctx context.Context, id int) (model.User, error) {
	query := m.db.Rebind(SelectQuery)
	row := m.db.QueryRowContext(ctx, query, id)
	var u model.User
	if err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Specialization, &u.DOB); err != nil {
		return model.User{}, errors.Wrap(err, "error scanning row")
	}
	return u, nil
}

func (m MariaDB) Update(ctx context.Context, u model.User) (model.User, error) {
	query := m.db.Rebind(UpdateQuery)
	res, err := m.db.ExecContext(ctx, query, u.FirstName, u.LastName, u.Specialization, u.DOB, u.ID)
	if err != nil {
		return model.User{}, errors.Wrap(err, "error updating data in user table")
	}

	if n, err := res.RowsAffected(); n == 0 {
		if err != nil {
			return model.User{}, errors.Wrap(err, "error reading res.RowsAffected() method")
		}
		return model.User{}, errors.Wrap(errors.New("nothing changed"), "error updating a row")
	}

	resp, err := m.User(ctx, int(u.ID))
	if err != nil {
		return model.User{}, errors.Wrap(err, "error parsing data right after its update")
	}
	return resp, err
}

func (m MariaDB) Delete(ctx context.Context, id int) error {
	query := m.db.Rebind(DeleteQuery)
	res, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "could not remove data")
	}

	if n, err := res.RowsAffected(); n == 0 {
		if err != nil {
			return errors.Wrap(err, "error reading res.RowsAffected() method")
		}
		return errors.Wrap(errors.New("given id does not exist"), "error deleting a row")
	}
	return nil
}
