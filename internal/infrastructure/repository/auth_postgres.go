package repository

import (
	"fmt"

	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func newAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres{
	return &AuthorizationPostgres{db:db}
}

func (au * AuthorizationPostgres) SignUp(input entity.RegiterInput) error{
	uid:=uuid.NewString()

	query:=fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE email=$1 or phone=$2", usersTable)

	var count int 
	row:= au.db.QueryRow(query,input.Email,input.Phone)
	if err:=row.Scan(&count);err!=nil{
		return err
	}
	
	if count>0{
		return entity.ErrUserAlredyExists
	}
	query= fmt.Sprintf("INSERT INTO %s (id,name,phone,email) VALUES($1,$2,$3,$3)",usersTable)
	_,err:=au.db.Exec(query,uid,input.Name,input.Phone,input.Email)
	if err!=nil{
		return err
	}
	return nil
}