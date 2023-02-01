package repository

import (
	"time"

	"github.com/C-STYR/go-web1/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(r models.RoomRestriction) error

	SearchAvailabilityByDatesByRoomId(start, end time.Time, roomID int) (bool, error)
	
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
