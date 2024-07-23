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

func newAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (au *AuthorizationPostgres) Register(input entity.RegiterInput) error {
	uid := uuid.NewString()

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE email=$1 or phone=$2", usersTable)

	var count int
	row := au.db.QueryRow(query, input.Email, input.Phone)
	if err := row.Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		return entity.ErrUserAlredyExists
	}
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE email=$1 or phone=$2", usersTmpTable)
	var countTmp int
	row = au.db.QueryRow(query, input.Email, input.Phone)
	if err := row.Scan(&countTmp); err != nil {
		return err
	}

	if countTmp > 0 {
		query := fmt.Sprintf("UPDATE %s SET uid=&1 WHERE email=$2 or phone=$3", usersTmpTable)
		_, err := au.db.Exec(query, uid, input.Email,input.Phone)
		if err != nil {
			return err
		}
	} else {

		query = fmt.Sprintf("INSERT INTO %s (id,name,phone,email,birthdate) VALUES($1,$2,$3,$4,$5)", usersTmpTable)
		_, err := au.db.Exec(query, uid, input.Name, input.Phone, input.Email,input.Birthdate)
		if err != nil {
			return err
		}
	}
	return nil
}

func (au *AuthorizationPostgres) SendCodeEmail(email string, code string) error{
	query:=fmt.Sprintf("UPDATE %s SET code=$1 WHERE email=$1",usersTmpTable)
	_, err:= au.db.Exec(query,code,email)
	if err!=nil{
		return err
	}
	return nil
}

func (au *AuthorizationPostgres) SendCodeSMS(phone string, code string) error{
	query:=fmt.Sprintf("UPDATE %s SET code=$1 WHERE phone=$1",usersTmpTable)
	_, err:= au.db.Exec(query,code,phone)
	if err!=nil{
		return err
	}
	return nil
}