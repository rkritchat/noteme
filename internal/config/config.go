package config

import (
	"fmt"
	"github.com/gobuffalo/pop"
	"os"
)

type config struct {
	DB *pop.Connection
}

func NewConfig() *config {
	cfg := new(config)
	cfg.initDB()
	return cfg
}

func (c *config) initDB() {
	connection, err := pop.NewConnection(&pop.ConnectionDetails{
		Driver: "mysql",
		URL:    fmt.Sprintf("mysql://%v:%v@tcp(%v)/%v?parseTime=true&multiStatements=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PWD"), os.Getenv("MYSQL_ADDR"), os.Getenv("MYSQL_DB_NAME")),
	})
	if err != nil {
		panic(err)
	}

	err = connection.Open()
	if err != nil {
		panic(err)
	}
	c.DB = connection
}
