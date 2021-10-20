package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/api"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	cfg, err := configs.EnvParser()
	if err != nil {
		log.Fatalln(err.Error())
	}

	mariaDB, err := mariadb.Connect(cfg.MariaDB)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = mariaDB.Up(); err != nil {
		log.Fatalln(err.Error())
	}

	srv := service.Conn{DB: mariaDB}
	r := api.Router{Service: srv}

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(http.ListenAndServe(addr, api.Handlers(r)))
}
