package configs

type Config struct {
	Server  Server  `yaml:"server"`
	MariaDb MariaDb `yaml:"mariadb"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MariaDb struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}
