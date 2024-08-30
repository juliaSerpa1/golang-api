package infra

import (
    "errors"
    "sync"
    "golang-api/domain"
)

type TestEntityRepository struct {
    mu    sync.Mutex
    store map[int]domain.TestEntity
    nextID int
}

func NewTestEntityRepository() *TestEntityRepository {
    return &TestEntityRepository{
        store: make(map[int]domain.TestEntity),
        nextID: 1,
    }
}

func (r *TestEntityRepository) Create(entity domain.TestEntity) domain.TestEntity {
    r.mu.Lock()
    defer r.mu.Unlock()

    entity.ID = r.nextID
    r.nextID++
    r.store[entity.ID] = entity
    return entity
}

func (r *TestEntityRepository) GetByID(id int) (domain.TestEntity, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    entity, exists := r.store[id]
    if !exists {
        return domain.TestEntity{}, errors.New("entity not found")
    }
    return entity, nil
}

func (r *TestEntityRepository) Update(entity domain.TestEntity) (domain.TestEntity, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    _, exists := r.store[entity.ID]
    if !exists {
        return domain.TestEntity{}, errors.New("entity not found")
    }
    r.store[entity.ID] = entity
    return entity, nil
}

func (r *TestEntityRepository) Delete(id int) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    _, exists := r.store[id]
    if !exists {
        return errors.New("entity not found")
    }
    delete(r.store, id)
    return nil
}