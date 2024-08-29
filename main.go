package main

import (
    "log"
    "golang-api/di"
    "golang-api/router"
)

func main() {
    // Inicializa as dependências
    container := di.Initialize()

    // Configura as rotas
    r := router.SetupRoutes(container)

    // Inicializa o servidor
    log.Fatal(r.Run(":8080"))
}