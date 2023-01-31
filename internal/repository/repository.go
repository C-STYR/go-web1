package repository

import "github.com/C-STYR/go-web1/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error 
}