package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Postgres struct{
	User string
	Host string
	Port string
	Dbname string
	Sslmode string
	Password string
}

// New 
func New(url Postgres) (*sqlx.DB, error) {
	db,err:= sqlx.Open("postgres",fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	url.Host, url.Port, url.User, url.Password, url.Dbname, url.Sslmode))
	if err!=nil{
		return nil,err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db,nil
}


