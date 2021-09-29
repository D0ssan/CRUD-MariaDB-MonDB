package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/d0ssan/CRUD-MariaDB-MongoDB/api"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"
	mariadb "github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/service"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	mariaDb, err := mariadb.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}

	srv := service.New(mariaDb)
	handler := api.New(srv)

	addrCfg := new(configs.Server)
	if err := envconfig.Process("myserver", addrCfg); err != nil {
		log.Fatal(err.Error())
	}
	addr := fmt.Sprintf("%v:%v", addrCfg.Host, addrCfg.Port)
	log.Fatal(http.ListenAndServe(addr, handler))
}
