package cinemaservice

import (
	"cinema/pkg/mpostgres"
	"cinema/pkg/utils"
	"errors"
)

type Repository interface {
	CreateCinema(name string, rows, columns int) (int, error)
	GetCinemaConfig(cinemaID int) (Cinema, error)
	ReserveSeats(cinemaID int, seats []Seat, minDistance int) error
	CancelSeats(cinemaID int, seats []Seat) error
	GetReservedSeats(cinemaID int) ([]Seat, error)
}

type repository struct{}

func newCinemaRepository() *repository {
	return &repository{}
}

// Creates a new cinema
func (repo *repository) CreateCinema(name string, rows, columns int) (int, error) {
	db := mpostgres.GetDB()
	var cinemaID int
	err := db.QueryRow(
		"INSERT INTO cinemas (name, rows, columns) VALUES ($1, $2, $3) RETURNING id",
		name, rows, columns,
	).Scan(&cinemaID)
	if err != nil {
		return 0, err
	}
	return cinemaID, nil
}

func (repo *repository) GetCinemaConfig(cinemaID int) (Cinema, error) {
	db := mpostgres.GetDB()

	var cinema Cinema

	err := db.QueryRow("SELECT name, rows, columns FROM cinemas WHERE id = $1", cinemaID).Scan(&cinema.Name, &cinema.Rows, &cinema.Columns)
	if err != nil {
		return cinema, err
	}
	return cinema, nil
}

// Reserves seats
func (repo *repository) ReserveSeats(cinemaID int, seats []Seat, minDistance int) error {
	db := mpostgres.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Validate distancing
	reservedSeats, err := repo.GetReservedSeats(cinemaID)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, seat := range seats {
		if violatesDistancing(seat, reservedSeats, minDistance) {
			tx.Rollback()
			return errors.New("reservation violates distancing rules")
		}

		_, err := tx.Exec(
			"INSERT INTO reserved_seats (cinema_id, row, col, status) VALUES ($1, $2, $3, 'reserved')",
			cinemaID, seat.Row, seat.Column,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// Cancels seats
func (repo *repository) CancelSeats(cinemaID int, seats []Seat) error {
	db := mpostgres.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, seat := range seats {
		_, err := tx.Exec(
			"UPDATE reserved_seats SET status = 'canceled' WHERE cinema_id = $1 AND row = $2 AND col = $3 AND status = 'reserved'",
			cinemaID, seat.Row, seat.Column,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// Fetch reserved seats for a cinema
func (repo *repository) GetReservedSeats(cinemaID int) ([]Seat, error) {
	db := mpostgres.GetDB()
	rows, err := db.Query("SELECT row, col FROM reserved_seats WHERE cinema_id = $1 AND status = 'reserved'", cinemaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservedSeats []Seat
	for rows.Next() {
		var seat Seat
		if err := rows.Scan(&seat.Row, &seat.Column); err != nil {
			return nil, err
		}
		reservedSeats = append(reservedSeats, seat)
	}
	return reservedSeats, nil
}

// Checks if a seat violates distancing rules
func violatesDistancing(seat Seat, reservedSeats []Seat, minDistance int) bool {
	for _, reserved := range reservedSeats {
		if utils.ManhattanDistance(seat.Row, seat.Column, reserved.Row, reserved.Column) < minDistance {
			return true
		}
	}
	return false
}
