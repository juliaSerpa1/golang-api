package mockjwt

import (
    "github.com/dgrijalva/jwt-go"
)

type MockJWTService struct{}

func NewMockJWTService() *MockJWTService {
    return &MockJWTService{}
}

func (m *MockJWTService) GenerateToken(userID int) (string, error) {
    return "mock-token", nil
}

func (m *MockJWTService) ValidateToken(tokenString string) (*jwt.Token, error)  {
    return &jwt.Token{
        Valid: true,
    }, nil
}