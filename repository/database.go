package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type DataSource struct {
	*sqlx.DB
	maxIdleConns     int
	maxOpenConns     int
	maxConnsLifetime time.Duration
}

func NewDatabaseConnectionWithConnectionPool(host string, port string, user string, password string, database string, maxIdle int, maxOpenConnection int) (*DataSource, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database)
	db, err := sqlx.Connect("mysql", url)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	maxLifeTime := 1 * time.Second

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetConnMaxLifetime(maxLifeTime)
	return &DataSource{db, maxIdle, maxOpenConnection, maxLifeTime}, nil
}

func NewDatabaseConnection(host string, port string, user string, password string, database string) (*DataSource, error) {
	return NewDatabaseConnectionWithConnectionPool(host, port, user, password, database, 10, 10)
}
