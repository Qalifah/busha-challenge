package config

type PostgresConfig struct {
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type BaseConfig struct {
	ServePort string          `yaml:"serve_port"`
	Postgres  *PostgresConfig `yaml:"postgres"`
	Redis     *Redis          `yaml:"redis"`
}
