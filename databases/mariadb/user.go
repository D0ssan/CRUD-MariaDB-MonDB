package mariadb

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/model"
)

const (
	insertUserQuery = `INSERT INTO users (first_name, last_name, specialization, dob) VALUES (?,?,?,?)`
	userQuery       = `SELECT * FROM users WHERE id = ?`
	updUserQuery    = `UPDATE users SET first_name=?, last_name=?, specialization=?, dob=? WHERE id= ?`
	rmvUserQuery    = `DELETE FROM users WHERE id = ?`
	allUserQuery    = `SELECT * FROM users`
)

var (
	// ErrNoChange indicates no changes in the db during insert, delete or update operations.
	ErrNoChange = errors.New("no change")

	// ErrNoID  indicates that the db does not have such id.
	ErrNoID = errors.New("id does not exist")
)

// Insert save a data into the db and modifies the input model.User ID by writing it from the db.
func (m MariaDB) Insert(ctx context.Context, u *model.User) error {
	query := insertUserQuery

	// comma after query is safe  https://blog.sqreen.com/preventing-sql-injections-in-go-and-other-vulnerabilities/
	res, err := m.db.ExecContext(ctx, query, u.FirstName, u.LastName, u.Specialization, u.DOB)
	if err != nil {
		return errors.Wrap(err, "error executing the insert query into user table")
	}

	var n int64
	if n, err = res.RowsAffected(); n == 0 {
		if err != nil {
			return errors.Wrap(err, "error reading res.RowsAffected() method")
		}

		return errors.Wrap(ErrNoChange, "error inserting data")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "error parsing the last id from user table")
	}

	u.ID = id

	return nil
}

// User return model.User based on id from the db.
func (m MariaDB) User(ctx context.Context, id int) (model.User, error) {
	query := userQuery
	row := m.db.QueryRowContext(ctx, query, id)

	var u model.User

	if err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Specialization, &u.DOB); err != nil {
		return model.User{}, errors.Wrap(err, "error scanning row")
	}

	return u, nil
}

// Update changes a row taken from input model.User based on its id.
// If there is not any changes, the function returns with error ErrNoChange.
func (m MariaDB) Update(ctx context.Context, u model.User) error {
	query := updUserQuery

	res, err := m.db.ExecContext(ctx, query, u.FirstName, u.LastName, u.Specialization, u.DOB, u.ID)
	if err != nil {
		return errors.Wrap(err, "error updating data in user table")
	}

	var n int64
	if n, err = res.RowsAffected(); n == 0 {
		if err != nil {
			return errors.Wrap(err, "error reading res.RowsAffected() method")
		}

		return errors.Wrap(ErrNoChange, "error updating a row")
	}

	return err
}

// Delete remove a row from the db based on id. If there is no such id, the function return error ErrNoID.
func (m MariaDB) Delete(ctx context.Context, id int) error {
	query := rmvUserQuery

	res, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "could not remove data")
	}

	var n int64
	if n, err = res.RowsAffected(); n == 0 {
		if err != nil {
			return errors.Wrap(err, "error reading res.RowsAffected() method")
		}

		return errors.Wrap(ErrNoID, "error deleting a row")
	}

	return nil
}

// All returns all rows in slice of model.User.
func (m MariaDB) All(ctx context.Context) ([]model.User, error) {
	query := allUserQuery

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "could not make query")
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}(rows)

	var (
		u     model.User
		users []model.User
	)

	for rows.Next() {
		if err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Specialization, &u.DOB); err != nil {
			return nil, errors.Wrap(err, "could not scan rows")
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows error")
	}

	return users, nil
}
