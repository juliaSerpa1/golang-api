package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "golang-api/domain"
)

type UserController struct {
    DBService db.DBService
}

func NewUserController(dbService db.DBService) *UserController {
    return &UserController{DBService: dbService}
}

func (c *UserController) GetUser(ctx *gin.Context) {
    userID := ctx.Param("id")

    // Lógica para obter o usuário do banco de dados
    user := domain.User{ID: 1, Username: "example"}
    ctx.JSON(http.StatusOK, user)
}