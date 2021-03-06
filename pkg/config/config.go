package config

import (
	"flag"
	"fmt"
	"os"
)

// Config holds an app db configuration
type Config struct {
	dbUser     string
	dbPswd     string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
}

// Get returns a new server config
func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", os.Getenv("POSTGRES_PASSWORD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "DB name")
	flag.StringVar(&conf.testDBHost, "testdbhost", os.Getenv("TEST_DB_HOST"), "test database host")
	flag.StringVar(&conf.testDBName, "testdbname", os.Getenv("TEST_DB_NAME"), "test database name")

	flag.Parse()

	return conf
}

// GetDBConnStr returns a db connection
func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

// GetTestDBConnStr returns a test db connection
func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.testDBHost, c.testDBName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPswd,
		dbhost,
		c.dbPort,
		dbname,
	)
}
