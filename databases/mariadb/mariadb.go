package mariadb

import (
	"fmt"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type MariaDB struct {
	db *sqlx.DB
}

func Connect() (*MariaDB, error) {
	var dbCfg configs.MariaDb
	if err := envconfig.Process("MYMARIADB", &dbCfg); err != nil {
		return nil, errors.Wrap(err, "error processing mariadb configs")
	}
	// "root:secret@tcp(localhost:3306)/users"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",dbCfg.Username,dbCfg.Password,dbCfg.Host,dbCfg.Port, dbCfg.Name)
	db, err := sqlx.Connect(dbCfg.Driver, dsn) // all-all in config
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to the mariaDb")
	}
	return &MariaDB{db: db}, nil
}

// docker pull mariadb
// docker run ....
// migrate -path databases/mariadb/migration -database "mysql://root:secret@tcp(localhost:3306)/users" up
