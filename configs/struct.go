package configs

type Config struct {
	Server  Server
	MariaDB MariaDB
}

type Server struct {
	Host string
	Port string
}

type MariaDB struct {
	Driver        string
	Username      string
	Name          string
	Host          string
	Port          string
	Password      string
	PathToMigrate string
}
