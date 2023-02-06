package dbrepo

import (
	"errors"
	"time"

	"github.com/C-STYR/go-web1/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1, nil
}

// InsertRoomRestriction inserts a room restriction in the DB
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDates returns true if the selected dates/room are available
func (m *testDBRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomID int) (bool, error) {
	return true, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID returns room information by room ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
} 

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}