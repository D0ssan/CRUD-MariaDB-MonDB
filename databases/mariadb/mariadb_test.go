package mariadb_test

import (
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/stretchr/testify/assert"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb"
)

func TestConnect(t *testing.T) {
	tt := []struct {
		name string
		cfg  configs.MariaDB
		err  string
	}{
		{
			"Success connection",
			configs.MariaDB{
				Driver:        "mysql",
				Username:      "root",
				Name:          "test_users",
				Host:          "localhost",
				Port:          "3306",
				Password:      "secret",
				PathToMigrate: "migration",
			},
			"",
		},
		{
			"Failed connection: wrong dns",
			configs.MariaDB{
				Driver:        "mysql",
				Username:      "root",
				Name:          "test_users",
				Host:          "localhost",
				Password:      "secret",
				PathToMigrate: "migration",
			},
			"could not connect to the database",
		},
		{
			"Failed connection: wrong driver",
			configs.MariaDB{
				Driver:        "postgres",
				Username:      "root",
				Name:          "test_users",
				Host:          "localhost",
				Port:          "3306",
				Password:      "secret",
				PathToMigrate: "databases/mariadb/migrations",
			},
			"could not create migration migration",
		},
		{
			"Failed connection: wrong path to migrate",
			configs.MariaDB{
				Driver:        "mysql",
				Username:      "root",
				Name:          "test_users",
				Host:          "localhost",
				Port:          "3306",
				Password:      "secret",
				PathToMigrate: "WRONG_PATH",
			},
			"error creating db migration",
		},
		{
			"Failed connection: cannot migrate up and down",
			configs.MariaDB{
				Driver:        "mysql",
				Username:      "root",
				Name:          "test_users",
				Host:          "localhost",
				Port:          "3306",
				Password:      "secret",
				PathToMigrate: "",
			},
			"cannot manipulate migration",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db, err := mariadb.Connect(tc.cfg)
			if tc.err != "" {
				if tc.err == "cannot manipulate migration" {
					assert.NotNil(t, db)
					assert.Error(t, db.Up())
					assert.Error(t, db.Down())
					return
				}
				assert.Nil(t, db)
				assert.Error(t, err, tc.err)
				return
			}
			assert.NotNil(t, db)
			assert.NoError(t, err)
			assert.NoError(t, db.Up())
			assert.NoError(t, db.Down())
		})
	}
}
