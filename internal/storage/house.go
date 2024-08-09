package storage

import (
	"avito_bootcamp/internal/entity"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (s *Storage) GetHouse(ctx context.Context, houseID int, userType string) ([]entity.Flat, error) {
	var query string
	if userType == "client" {
		query = `select * from flats where house_id = $1 where status = 'approved'`
	} else if userType == "moderator" {
		query = `select * from flats where house_id = $1`
	}
	rows, err := s.Db.Query(query, houseID)
	if err != nil {
		return nil, err
	}
	var flats []entity.Flat
	for rows.Next() {
		var flat entity.Flat
		if err = rows.Scan(&flat.HouseID, &flat.Number, &flat.Price, &flat.Rooms, &flat.Status, &flat.CreatedAt); err != nil {
			log.Printf("Internal Server Error, %d", http.StatusInternalServerError)
		}
		flats = append(flats, flat)
	}
	return flats, nil
}
func (s *Storage) CreateHouse(ctx context.Context, house entity.House) error {
	query := `insert into houses(id, address, year_built, developer, created_at) values ($1, $2, $3, $4, $5)`
	err := s.Db.QueryRow(query, house.ID, house.Address, house.YearBuilt, house.Developer, time.Now()).Err()
	if err != nil {
		return fmt.Errorf("internal Server Error, %d", http.StatusInternalServerError)
	}
	return nil
}
