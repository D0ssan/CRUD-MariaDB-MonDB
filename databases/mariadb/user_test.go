package mariadb_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/model"
)

func TestOne(t *testing.T) {
	ConfDB := configs.MariaDB{
		Driver:        "mysql",
		Username:      "root",
		Name:          "test_users",
		Host:          "localhost",
		Port:          "3306",
		Password:      "secret",
		PathToMigrate: "migration",
	}
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
		err      string
		expected model.User
		id       int
	}{
		{"Success a user parsing ", "", user, 1},
		{"Failed a user parsing: no id in the db", "no such row", model.User{}, 2},
		{"Failure a user parsing: the db is dropped", "error db is dropped", model.User{}, 1},
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
	ConfDB := configs.MariaDB{
		Driver:        "mysql",
		Username:      "root",
		Name:          "test_users",
		Host:          "localhost",
		Port:          "3306",
		Password:      "secret",
		PathToMigrate: "migration",
	}
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
		err      string
		toUpdate model.User
		expected model.User
	}{
		{
			"Failure a user update: exact same values",
			"error no changes",
			user,
			model.User{},
		},
		{
			"Success a user update",
			"",
			toUpdate,
			toUpdate,
		},
		{
			"Failure a user update: no such id",
			"error no this id in the db",
			model.User{
				ID:        3,
				FirstName: "UpdatedTestFirstName",
			},
			model.User{},
		},
		{
			"Failure a user update: no such id",
			"error do is dropped",
			toUpdate,
			model.User{},
		},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if i == len(tt)-1 {
				err = db.Down()
				require.NoError(t, err)
			}

			err = db.Update(context.Background(), tc.toUpdate)
			update, _ := db.User(context.Background(), int(tc.toUpdate.ID))
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
	ConfDB := configs.MariaDB{
		Driver:        "mysql",
		Username:      "root",
		Name:          "test_users",
		Host:          "localhost",
		Port:          "3306",
		Password:      "secret",
		PathToMigrate: "migration",
	}
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
		err  string
		id   int
	}{
		{"Failure a user delete: no such row", "error no this id in the db", 4},
		{"Success a user delete", "", 1},
		{"Failure a user delete: db is dropped", "error db is dropped", 1},
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
	ConfDB := configs.MariaDB{
		Driver:        "mysql",
		Username:      "root",
		Name:          "test_users",
		Host:          "localhost",
		Port:          "3306",
		Password:      "secret",
		PathToMigrate: "migration",
	}
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
		err      string
		expected []model.User
	}{
		{"Success parsing all data", "", []model.User{user}},
		{"Failure parsing all data: the db is dropped", "error the db is dropped", nil},
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
	ConfDB := configs.MariaDB{
		Driver:        "mysql",
		Username:      "root",
		Name:          "test_users",
		Host:          "localhost",
		Port:          "3306",
		Password:      "secret",
		PathToMigrate: "migration",
	}
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
		err  string
		user model.User
	}{
		{"Success parsing all data", "", user},
		{"Failure parsing all data: the db is dropped", "error the db is dropped", user},
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
