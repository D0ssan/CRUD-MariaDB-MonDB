package mariadb

import (
	"database/sql"
	"fmt"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"

	"github.com/pkg/errors"
)

type MariaDB struct {
	db      *sql.DB
	migrate *migration
}

func Connect(cfg configs.MariaDB) (*MariaDB, error) {
	// "root:secret@tcp(localhost:3306)/users"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to the mariaDb")
	}

	if err := db.Ping(); err != nil {
		if err := db.Close(); err != nil {
			return nil, errors.Wrap(err, "error closing the database")
		}
		return nil, errors.Wrap(err, "error ping to the mariadb")
	}

	m, err := migrationInit(db, cfg.PathToMigrate, cfg.Driver)
	if err != nil {
		return nil, errors.Wrap(err, "error creating migrate migration")
	}

	return &MariaDB{db: db, migrate: m}, nil
}

// docker pull mariadb
// docker run exec -it <DOCKER ID> /bin/bash
// mysql -u root -p -h localhost
// CREATE <database>;
// migrate -path databases/mariadb/Migration -database "mysql://root:ppasword@tcp(localhost:3306)/users" up
