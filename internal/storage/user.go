package storage

import (
	"avito_bootcamp/internal/entity"
	"context"
	"fmt"
	"net/http"
	"time"
)

func (s *Storage) Login(ctx context.Context, user entity.User) (entity.User, error) {
	query := `select * from users where email = $1 and password = $2`

	rows, err := s.Db.Query(query, user.Email, user.Password)
	defer rows.Close()
	if err != nil {
		return entity.User{}, fmt.Errorf("internal server error, %d", http.StatusInternalServerError)
	}
	if !rows.Next() {
		return entity.User{}, fmt.Errorf("no such email or wrong password, %d", http.StatusInternalServerError)
	}
	var t time.Time
	err = rows.Scan(&user.Email, &user.Password, &user.UserType, &t)
	if err != nil {
		return entity.User{}, fmt.Errorf("internal server error, %d", http.StatusInternalServerError)
	}
	return user, nil
}

func (s *Storage) Register(ctx context.Context, user entity.User) error {
	query := `insert into users (email, password, user_type, created_at) values ($1, $2, $3, $4)`
	err := s.Db.QueryRow(query, user.Email, user.Password, user.UserType, time.Now()).Err()
	if err != nil {
		return fmt.Errorf("internal Server Error, %d", http.StatusInternalServerError)
	}
	return nil
}
