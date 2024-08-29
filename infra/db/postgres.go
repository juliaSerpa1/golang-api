package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

type DBService interface {
    // Métodos de acesso ao banco de dados
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