package mariadb

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/pkg/errors"
)

type migration struct {
	instance *migrate.Migrate
}

func migrationInit(db *sql.DB, path, cfgDriver string) (*migration, error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "error getting the driver")
	}

	instanceDB, err := migrate.NewWithDatabaseInstance("file://" + path, cfgDriver, driver)
	if err != nil {
		return nil, errors.Wrap(err, "error getting the database migration")
	}

	m := migration{instance: instanceDB}
	return &m, nil
}

func (m MariaDB) Up() error {
	err := m.migrate.instance.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return errors.Wrap(err, "could not migrate up")
	}
	return nil
}

func (m MariaDB) Down() error {
	if err := m.migrate.instance.Down(); err != nil {
		return errors.Wrap(err, "could not migrate down")
	}
	return nil
}
