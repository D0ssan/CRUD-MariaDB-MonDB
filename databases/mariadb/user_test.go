package mariadb_test

import (
	"context"
	"testing"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOne(t *testing.T) {
	db, err := mariadb.Connect(ConfDB)
	require.NoError(t, err)
	require.NotNil(t, db)
	err = db.Up()
	require.NoError(t, err)

	user := model.User{
		ID:             1,
		FirstName:      "TestFirstName",
		LastName:       "TestLastName",
		Specialization: "No one",
		DOB:            "2000-01-01",
	}

	err = db.Insert(context.Background(), &user)
	require.NoError(t, err)

	tt := []struct {
		name     string
		id       int
		expected model.User
		err      string
	}{
		{"Success a user parsing ", 1, user, ""},
		{"Failed a user parsing: no id in the db", 2, model.User{}, "no such row"},
		{"Failure a user parsing: the db is dropped", 1, model.User{}, "error db is dropped"},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if i == len(tt)-1 {
				err = db.Down()
				require.NoError(t, err)
			}
			u, err := db.User(context.Background(), tc.id)
			if tc.err != "" {
				assert.Error(t, err)
				assert.Equal(t, model.User{}, u)

				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, u)
		})
	}
}

func TestUpdate(t *testing.T) {
	db, err := mariadb.Connect(ConfDB)
	require.NoError(t, err)
	require.NotNil(t, db)
	err = db.Up()
	require.NoError(t, err)

	user := model.User{
		ID:             1,
		FirstName:      "TestFirstName",
		LastName:       "TestLastName",
		Specialization: "No one",
		DOB:            "2000-01-01",
	}

	err = db.Insert(context.Background(), &user)
	require.NoError(t, err)

	toUpdate := model.User{
		ID:             1,
		FirstName:      "UpdatedTestFirstName",
		LastName:       "UpdatedTestLastName",
		Specialization: "Someone",
		DOB:            "2001-01-01",
	}

	tt := []struct {
		name     string
		toUpdate model.User
		expected model.User
		err      string
	}{
		{
			"Failure a user update: exact same values",
			user,
			model.User{},
			"error no changes",
		},
		{
			"Success a user update",
			toUpdate,
			toUpdate,
			"",
		},
		{
			"Failure a user update: no such id",
			model.User{
				ID:        3,
				FirstName: "UpdatedTestFirstName",
			},
			model.User{},
			"error no this id in the db",
		},
		{
			"Failure a user update: no such id",
			toUpdate,
			model.User{},
			"error do is dropped",
		},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if i == len(tt)-1 {
				err = db.Down()
				require.NoError(t, err)
			}

			update, err := db.Update(context.Background(), tc.toUpdate)
			if tc.err != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expected, update)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, update)
		})
	}
}

func TestDelete(t *testing.T) {
	db, err := mariadb.Connect(ConfDB)
	require.NoError(t, err)
	require.NotNil(t, db)
	err = db.Up()
	require.NoError(t, err)

	user := model.User{
		ID:             1,
		FirstName:      "TestFirstName",
		LastName:       "TestLastName",
		Specialization: "No one",
		DOB:            "2000-01-01",
	}

	err = db.Insert(context.Background(), &user)
	require.NoError(t, err)

	tt := []struct {
		name string
		id   int
		err  string
	}{
		{"Failure a user delete: no such row", 4, "error no this id in the db"},
		{"Success a user delete", 1, ""},
		{"Failure a user delete: db is dropped", 1, "error db is dropped"},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if i == len(tt)-1 {
				err = db.Down()
				require.NoError(t, err)
			}

			err := db.Delete(context.Background(), tc.id)
			if tc.err != "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestAll(t *testing.T) {
	db, err := mariadb.Connect(ConfDB)
	require.NoError(t, err)
	require.NotNil(t, db)

	err = db.Up()
	require.NoError(t, err)

	user := model.User{
		ID:             1,
		FirstName:      "TestFirstName",
		LastName:       "TestLastName",
		Specialization: "No one",
		DOB:            "2000-01-01",
	}

	err = db.Insert(context.Background(), &user)
	require.NoError(t, err)

	tt := []struct {
		name     string
		expected []model.User
		err      string
	}{
		{"Success parsing all data", []model.User{user}, ""},
		{"Failure parsing all data: the db is dropped", nil, "error the db is dropped"},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if i == len(tt)-1 {
				err = db.Down()
				require.NoError(t, err)
			}
			users, err := db.All(context.Background())
			if tc.err != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expected, users)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, users)
		})
	}
}

func TestInsert(t *testing.T) {
	db, err := mariadb.Connect(ConfDB)
	require.NoError(t, err)
	require.NotNil(t, db)

	err = db.Up()
	require.NoError(t, err)

	user := model.User{
		ID:             1,
		FirstName:      "TestFirstName",
		LastName:       "TestLastName",
		Specialization: "No one",
		DOB:            "2000-01-01",
	}

	tt := []struct {
		name string
		user model.User
		err  string
	}{
		{"Success parsing all data", user, ""},
		{"Failure parsing all data: the db is dropped", user, "error the db is dropped"},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if i == len(tt)-1 {
				err = db.Down()
				require.NoError(t, err)
			}
			err := db.Insert(context.Background(), &tc.user)
			if tc.err != "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			parsedUser, err := db.User(context.Background(), int(tc.user.ID))
			assert.NoError(t, err)
			assert.Equal(t, parsedUser, user)
		})
	}
}
