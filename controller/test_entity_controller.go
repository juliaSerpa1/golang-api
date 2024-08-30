package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "golang-api/domain"
    "golang-api/infra"
)

type TestEntityController struct {
    repo *infra.TestEntityRepository
}

func NewTestEntityController(repo *infra.TestEntityRepository) *TestEntityController {
    return &TestEntityController{repo: repo}
}

func (c *TestEntityController) Create(w http.ResponseWriter, r *http.Request) {
    var entity domain.TestEntity
    if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    createdEntity := c.repo.Create(entity)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdEntity)
}

func (c *TestEntityController) GetByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    entity, err := c.repo.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(entity)
}

func (c *TestEntityController) Update(w http.ResponseWriter, r *http.Request) {
    var entity domain.TestEntity
    if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    updatedEntity, err := c.repo.Update(entity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(updatedEntity)
}

func (c *TestEntityController) Delete(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    err = c.repo.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}