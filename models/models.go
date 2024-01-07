package models

import "database/sql"

// Master struct for all models to be attached to
type Models struct{}

func NewModels(DB *sql.DB) *Models {
	return &Models{}
}
