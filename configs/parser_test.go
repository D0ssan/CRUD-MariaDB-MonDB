package configs_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"
)

// TestEnvParser test EnvParser() func.
func TestEnvParser(t *testing.T) {
	tt := []struct {
		name     string
		expected configs.Config
		err      string
	}{
		{
			"Only MYMARIADB_PASSWORD",
			configs.Config{
				Server: configs.Server{},
				MariaDB: configs.MariaDB{
					Password: os.Getenv("MYMARIADB_PASSWORD"),
				},
			},
			"",
		},
		{
			"all MYMARIDB configs",
			configs.Config{
				Server: configs.Server{},
				MariaDB: configs.MariaDB{
					Driver:        "mysql",
					Username:      "root",
					Name:          "test_users",
					Host:          "localhost",
					Port:          "3306",
					Password:      os.Getenv("MYMARIADB_PASSWORD"),
					PathToMigrate: "file://",
				},
			},
			"",
		},
		{
			"all MYMARIDB configs plus ONE any",
			configs.Config{
				Server: configs.Server{},
				MariaDB: configs.MariaDB{
					Driver:        "mysql",
					Username:      "root",
					Name:          "test_users",
					Host:          "localhost",
					Port:          "3306",
					Password:      os.Getenv("MYMARIADB_PASSWORD"),
					PathToMigrate: "file://",
				},
			},
			"",
		},
		{
			"all MYMARIADB and all MYSERVER",
			configs.Config{
				Server: configs.Server{
					Host: "localhost",
					Port: "8080",
				},
				MariaDB: configs.MariaDB{
					Driver:        "mysql",
					Username:      "root",
					Name:          "test_users",
					Host:          "localhost",
					Port:          "3306",
					Password:      os.Getenv("MYMARIADB_PASSWORD"),
					PathToMigrate: "file://",
				},
			},
			"",
		},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := configs.EnvParser()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, cfg)
		})

		switch i {
		case 0:
			err := os.Setenv("MYMARIADB_DRIVER", "mysql")
			require.NoError(t, err)

			err = os.Setenv("MYMARIADB_USERNAME", "root")
			require.NoError(t, err)

			err = os.Setenv("MYMARIADB_NAME", "test_users")
			require.NoError(t, err)

			err = os.Setenv("MYMARIADB_HOST", "localhost")
			require.NoError(t, err)

			err = os.Setenv("MYMARIADB_PORT", "3306")
			require.NoError(t, err)

			err = os.Setenv("MYMARIADB_PATHTOMIGRATE", "file://")
			require.NoError(t, err)

		case 1:
			err := os.Setenv("MYMARIADB_ANYTHING", "anything")
			require.NoError(t, err)
		case 2:
			err := os.Setenv("MYSERVER_HOST", "localhost")
			require.NoError(t, err)

			err = os.Setenv("MYSERVER_PORT", "8080")
			require.NoError(t, err)
		}
	}
}
