package main

import (
	"flag"
	"fmt"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/api"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/configs"
	mariadb "github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb"
	"github.com/d0ssan/CRUD-MariaDB-MongoDB/service"
	"log"
	"net/http"
)

func main() {
	cfgPath := envs()
	cfg, err := configs.New(cfgPath)
	if err != nil {
		log.Fatalf("error parsing a .yml file: %v", err.Error())
	}

	mariaDb, err := mariadb.Connect(cfg.MariaDb)
	if err != nil {
		log.Fatalln(err.Error())
	}

	srv := service.New(mariaDb)
	handler := api.New(srv)

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(http.ListenAndServe(addr, handler))
}

// envs reads command line flags for yaml connection
func envs() string {
	var cfgPath string
	flag.StringVar(
		&cfgPath,
		"configPath",
		"configs/config.yml",
		"-configPath flag is reference to a .yaml file, example: -configPath=configs/config.yml",
	)
	flag.Parse()
	return cfgPath
}
