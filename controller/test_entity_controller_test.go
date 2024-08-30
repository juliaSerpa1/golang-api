package controller

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "golang-api/domain"
    "golang-api/infra"
)

func TestCreate(t *testing.T) {
    repo := infra.NewTestEntityRepository()
    ctrl := NewTestEntityController(repo)

    entity := domain.TestEntity{Name: "Test", Value: "Value"}
    body, _ := json.Marshal(entity)
    req := httptest.NewRequest(http.MethodPost, "/test-entity", bytes.NewReader(body))
    w := httptest.NewRecorder()

    ctrl.Create(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusCreated {
        t.Fatalf("expected status %v, got %v", http.StatusCreated, res.StatusCode)
    }

    var createdEntity domain.TestEntity
    if err := json.NewDecoder(res.Body).Decode(&createdEntity); err != nil {
        t.Fatal(err)
    }
    if createdEntity.Name != entity.Name || createdEntity.Value != entity.Value {
        t.Errorf("expected %v, got %v", entity, createdEntity)
    }
}