package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // To provide mysql hook for db package
)

var db *sql.DB

var dbConfig DBConfig = DBConfig{
	Driver:   "mysql",
	User:     "root",
	Password: "admin@123",
	DBName:   "sys",
}

// GetConnection provides open database connection
func GetConnection() *sql.DB {
	db, err := sql.Open(dbConfig.Driver, dbConfig.GetMySQLConnStr())

	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}
