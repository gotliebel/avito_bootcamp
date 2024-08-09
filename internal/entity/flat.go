package entity

import (
	"time"
)

type Flat struct {
	HouseID   int       `db:"house_id" json:"house_id"`
	Number    int       `db:"number" json:"number"`
	Price     int       `db:"price" json:"price"`
	Rooms     int       `db:"rooms" json:"rooms"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
