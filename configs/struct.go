package configs

type Server struct {
	Host string
	Port string
}

type MariaDb struct {
	Driver   string
	Username string
	Name     string
	Host     string
	Port     string
	Password string
}
