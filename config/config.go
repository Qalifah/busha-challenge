package config

type PostgresConfig struct {
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	MaxConn  int    `yaml:"max_conn"`
}

type BaseConfig struct {
	ServePort string          `yaml:"serve_port"`
	Postgres  *PostgresConfig `yaml:"postgres"`
}
