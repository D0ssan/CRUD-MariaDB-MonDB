package configs

// Config has instances of Server and MariaDB configs.
type Config struct {
	Server  Server
	MariaDB MariaDB
}

// Server has host and port addresses.
type Server struct {
	Host string
	Port string
}

// MariaDB keeps information about the db instances and path to migration files.
type MariaDB struct {
	Driver        string
	Username      string
	Name          string
	Host          string
	Port          string
	Password      string
	PathToMigrate string
}
