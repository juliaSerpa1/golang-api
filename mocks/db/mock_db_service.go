package mockdb

import (
	"errors"
    "strconv"
	"golang-api/domain"
)

type MockDBService struct{}

func NewMockDBService() *MockDBService {
    return &MockDBService{}
}

func (m *MockDBService) GetUserByID(id string) (domain.User, error) {
    // Simule o comportamento desejado
    switch id {
    case "error":
        return domain.User{}, errSimulated
    default:
        // Converta o ID para int se necessário
        userID, err := strconv.Atoi(id)
        if err != nil {
            return domain.User{}, err
        }
        return domain.User{
            ID:       userID,
            Username: "mock-username",
        }, nil
    }
}

var errSimulated = errors.New("simulated error")