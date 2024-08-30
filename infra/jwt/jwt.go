package jwt

import (
    "github.com/dgrijalva/jwt-go"
    "os"
    "time"
)

var secretKey string

func InitializeJWT() {
    secretKey = getSecretKey()
}

func getSecretKey() string {
    key := os.Getenv("JWT_SECRET_KEY")
    if key == "" {
        panic("JWT_SECRET_KEY environment variable is not set")
    }
    return key
}

type JWTService interface {
    GenerateToken(userID int) (string, error)
    ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct{}

func NewJWTService() JWTService {
    return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 1).Unix(),
    })

    return token.SignedString([]byte(secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })
}