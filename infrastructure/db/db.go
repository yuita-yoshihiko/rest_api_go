package db

import (
	"database/sql"
	"fmt"
	"os"
	"rest_api_go/models"
	"time"

	"github.com/lib/pq"
)

var DB *sql.DB
var txKey = struct{}{}

type (
	DBAdministrator interface {
		Error(err error) error
		Queries() *models.Queries
	}
	dbutils struct {
		db      *sql.DB
		queries *models.Queries
	}
)

func NewDBAdministrator(db *sql.DB) DBAdministrator {
	return &dbutils{
		db:      db,
		queries: models.New(db),
	}
}

func Init() {
	connection, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	connection += " password=" + os.Getenv("DATABASE_PASSWORD") + " sslmode=" + os.Getenv("DATABASE_SSL_MODE")

	DB, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxLifetime(300 * time.Second)
}

func (d *dbutils) Error(err error) error {
	if err == nil {
		return nil
	}
	return err
}

func (d *dbutils) Queries() *models.Queries {
	return d.queries
}
