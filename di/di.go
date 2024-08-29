package di

import (
    "golang-api/infra/jwt"
    "golang-api/infra/db"
)

type Container struct {
    JWTService jwt.JWTService
    DBService  db.DBService
}

func Initialize() *Container {
    return &Container{
        JWTService: jwt.NewJWTService(),
        DBService:  db.NewDBService(),
    }
}