package entity

import (
	"time"
)

type House struct {
	ID              int       `db:"id" json:"id"`
	Address         string    `db:"address" json:"address"`
	YearBuilt       int       `db:"year_built" json:"year_built"`
	Developer       string    `db:"developer" json:"developer"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	LastFlatAddedAt time.Time `db:"last_flat_added_at" json:"last_flat_added_at"`
}
