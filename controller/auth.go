package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "golang-api/infra/jwt"
)

type AuthController struct {
    JWTService jwt.JWTService
}

func NewAuthController(jwtService jwt.JWTService) *AuthController {
    return &AuthController{JWTService: jwtService}
}

func (c *AuthController) Login(ctx *gin.Context) {
    var loginRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := ctx.BindJSON(&loginRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Lógica de autenticação e geração de token
    token, err := c.JWTService.GenerateToken(1) // Exemplo com ID fictício
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
}