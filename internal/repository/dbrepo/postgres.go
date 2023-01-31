package dbrepo

import (
	"context"
	"time"

	"github.com/C-STYR/go-web1/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}
// InsertReservation handles the posting of a reservation
func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	// this cancels our transaction if it isn't completed with 3 seconds (to account for service loss, etc)
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
		end_date, room_id, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
