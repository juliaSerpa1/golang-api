package router

import (
    "github.com/gin-gonic/gin"
    "golang-api/controller"
    "golang-api/di"
)

func SetupRoutes(container *di.Container) *gin.Engine {
    r := gin.Default()

    authController := controller.NewAuthController(container.JWTService)
    userController := controller.NewUserController(container.DBService)

    r.POST("/login", authController.Login)
    r.GET("/user/:id", userController.GetUser)

    return r
}