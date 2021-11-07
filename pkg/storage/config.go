package storage

import "fmt"

type DriverType int

const (
	PostgresSQL DriverType = iota + 1
	MongoDB
)

func (t DriverType) String() string {
	return [...]string{"PostgresSQL", "MongoDB"}[t-1]
}

type Config struct {
	host      string
	user      string
	password  string
	database  string
	port      int
	sslModeOn bool

	DriverType DriverType
}

func NewConfig(host, user, password, database string, port int, sslModeOn bool, DriverType DriverType) *Config {
	switch DriverType {
	case PostgresSQL:
		return &Config{
			host:       host,
			port:       port,
			user:       user,
			password:   password,
			database:   database,
			sslModeOn:  sslModeOn,
			DriverType: DriverType,
		}
	case MongoDB:
		return &Config{
			host:       host,
			port:       port,
			user:       user,
			password:   password,
			database:   database,
			sslModeOn:  sslModeOn,
			DriverType: DriverType,
		}
	}

	return nil
}

func (c *Config) ConnectionString() string {
	switch c.DriverType {
	case PostgresSQL:
		sslMode := "disabled"
		if c.sslModeOn {
			sslMode = "enabled"
		}
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			c.host, c.port, c.user, c.password, c.database, sslMode)
	case MongoDB:
		return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			c.user, c.password, c.host, c.port, c.database)
	}

	return ""
}
