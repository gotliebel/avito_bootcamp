package storage

import (
	"avito_bootcamp/internal/entity"
	"context"
	"fmt"
	"net/http"
	"time"
)

func (s *Storage) CreateFlat(ctx context.Context, flat entity.Flat) error {
	flat.Status = "created"
	curTime := time.Now()
	query := `insert into flats(house_id, number, price, rooms, status, created_at) values ($1, $2, $3, $4, $5, $6)`
	err := s.Db.QueryRow(query, flat.HouseID, flat.Number, flat.Price, flat.Rooms, flat.Status, curTime).Err()
	if err != nil {
		return fmt.Errorf("internal Server Error, %d", http.StatusInternalServerError)
	}
	query = `update houses set last_flat_added = $1 where id = $2`
	err = s.Db.QueryRow(query, curTime, flat.HouseID).Err()
	if err != nil {
		return fmt.Errorf("internal Server Error, %d", http.StatusInternalServerError)
	}
	return nil
}

func (s *Storage) UpdateFlat(ctx context.Context, flat entity.Flat) error {
	flat.Status = "approved"
	query := `update flats set status = $1 where house_id = $2 and number = $3`
	err := s.Db.QueryRow(query, flat.Status, flat.HouseID, flat.Number).Err()
	if err != nil {
		return fmt.Errorf("internal Server Error, %d", http.StatusInternalServerError)
	}
	return nil
}
