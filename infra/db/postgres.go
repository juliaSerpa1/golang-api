package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang-api/domain"
	// "log"
)

type DBService interface {
	GetUserByID(id string) (domain.User, error)
}

type dbService struct {
	db *sql.DB
}

func NewDBService() DBService {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return &dbService{db: db}
}

func (s *dbService) GetUserByID(id string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, username FROM users WHERE id=$1"
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}