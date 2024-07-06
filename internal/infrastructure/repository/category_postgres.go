package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/jmoiron/sqlx"
)


type CategoryPostgres struct {
	db *sqlx.DB
}

func newCategoryPostgres(db *sqlx.DB) *CategoryPostgres{
	return &CategoryPostgres{db:db}
}

func (cp *CategoryPostgres) CreateCategory(input entity.Category) (int, error){
	var categoryID int
	var row *sql.Row
	query:=fmt.Sprintf("INSERT INTO %s (name,parentid) values ($1,$2) RETURNING id",categoryTable)
	if input.ParentID==0{
		row=cp.db.QueryRow(query,input.Name,nil)
	} else{
		row=cp.db.QueryRow(query,input.Name,input.ParentID)
	}
	if err:=row.Scan(&categoryID);err!=nil{
		return 0,err
	}
	return categoryID,nil
}

func (cp *CategoryPostgres) GetAllCategorys() ([]entity.CategoryResult, error){
	var categories []entity.CategoryResult
	query:=fmt.Sprintf("SELECT t1.name, t2.name, t1.id FROM %s as t1 RIGHT JOIN %s AS t2 ON t2.parentid = t1.id", categoryTable,categoryTable)
	rows,err:=cp.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var result entity.CategoryResult
		if err:=rows.Scan(&result.Parentname,&result.ChildName,&result.ID);err!=nil{
			return nil,err
		}
		categories=append(categories, result)
	}
	return categories,nil
}

func (cp *CategoryPostgres) GetCategoryById(id int) (entity.CategoryResult, error){
	var  result entity.CategoryResult
	query:=fmt.Sprintf("SELECT t1.name, t2.name,t1.id FROM %s as 1 RIGHT JOIN %s AS t2 ON t2.parentid = t1.id WHERE t1.id=$1", categoryTable,categoryTable)
	row:=cp.db.QueryRow(query,id)
	err:=row.Scan(&result.Parentname,&result.ChildName,&result.ID)
	if err!=nil{
		return entity.CategoryResult{},err
	}
	return result, nil
}

func (cp *CategoryPostgres) UpdateCategory(id int, input entity.UpdateCategory) error{
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name!= nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.ParentID != nil {
		setValues = append(setValues, fmt.Sprintf("parentid=$%d", argId))
		args = append(args, *input.ParentID)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", categoryTable, setQuery, argId)

	args = append(args, id)

	_, err := cp.db.Exec(query, args...)
	if err!=nil{
		return err
	}

	return nil
}

func (cp *CategoryPostgres) DeleteCategory(id int) error{ 
	query:=fmt.Sprintf("DELETE FROM %s WHERE id=$1",categoryTable)
	_,err:=cp.db.Exec(query,id)
	if err!=nil{
		return err
	}
	return nil
}