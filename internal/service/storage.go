package service

import (
	"avito_bootcamp/internal/entity"
	"context"
)

type storage interface {
	Login(ctx context.Context, user entity.User) (entity.User, error)
	Register(ctx context.Context, user entity.User) error
	CreateFlat(ctx context.Context, flat entity.Flat) error
	UpdateFlat(ctx context.Context, flat entity.Flat) error
	GetHouse(ctx context.Context, houseID int, userType string) ([]entity.Flat, error)
	CreateHouse(ctx context.Context, house entity.House) error
}
