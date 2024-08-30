package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-api/controller"
	"golang-api/infra/db"
	"golang-api/infra/jwt"
    "golang-api/mocks/db"
    "golang-api/mocks/jwt"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockContainer struct {
    JWTService jwt.JWTService
    DBService  db.DBService
}

func SetupRoutes(jwtService jwt.JWTService, dbService db.DBService) *gin.Engine {
    router := gin.Default()
    authController := controller.NewAuthController(jwtService)
    userController := controller.NewUserController(dbService)

    router.POST("/login", authController.Login)
    router.GET("/user/:id", userController.GetUser)
    return router
}

func TestSetupRoutes(t *testing.T) {
    mockJWTService := mockjwt.NewMockJWTService()
    mockDBService := mockdb.NewMockDBService()

    router := SetupRoutes(mockJWTService, mockDBService)

    req, _ := http.NewRequest("POST", "/login", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)

    req, _ = http.NewRequest("GET", "/user/123", nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}