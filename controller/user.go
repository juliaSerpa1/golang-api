package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "golang-api/domain"
	"golang-api/infra/db"
)

type UserController struct {
	DBService db.DBService
}

func NewUserController(dbService db.DBService) *UserController {
	return &UserController{DBService: dbService}
}

func (c *UserController) GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := c.DBService.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}