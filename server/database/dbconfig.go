package database

import "fmt"

// DBConfig holds important information about the database configuration
type DBConfig struct {
	Driver   string
	User     string
	Password string
	DBName   string
}

// GetMySQLConnStr returns connection string to open MySql database
func (c *DBConfig) GetMySQLConnStr() string {
	return fmt.Sprintf("%s:%s@/%s", c.User, c.Password, c.DBName)
}
