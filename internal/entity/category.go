package entity

import "errors"

type Category struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	ParentID int    `json:"parent-id" default:"null"`
}

type UpdateCategory struct {
	Name     *string `json:"name"`
	ParentID *int    `json:"parent-id" default:"null"`
}

func Validate(input UpdateCategory) error {
	if input.Name == nil && input.ParentID == nil {
		return errors.New("Update category no validate")
	}
	return nil
}
