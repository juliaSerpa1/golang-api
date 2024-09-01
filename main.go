package main

import (
	"fmt"
	"golang-api/di"
	"golang-api/infra/jwt"
	"golang-api/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Inicializa o JWT
    jwt.InitializeJWT()

    // Acessar a variável de ambiente
    jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
    if jwtSecretKey == "" {
        log.Fatal("JWT_SECRET_KEY not set in environment")
    }

    // Imprimir a chave secreta para verificação (remova isso em produção)
    fmt.Println("JWT Secret Key:", jwtSecretKey)
    
    // Inicializa as dependências
    container := di.Initialize()

    // Supondo que container tenha métodos para obter os serviços
    jwtService := container.JWTService
    dbService := container.DBService

    // Configura as rotas
    r := router.SetupRoutes(jwtService, dbService)

    // Inicializa o servidor
    log.Fatal(r.Run(":8080"))
}