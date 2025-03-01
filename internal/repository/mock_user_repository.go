package repository

import (
	"errors"
	"sync"
	"trading-go/internal/model"
)

type MockUserRepository struct {
	data map[int]*model.User
	mu   sync.Mutex
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		data: make(map[int]*model.User),
	}
}

func (r *MockUserRepository) GetUserByID(id int) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.data[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *MockUserRepository) CreateUser(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.data[user.ID] = user
	return nil
}

func (r *MockUserRepository) DeleteUser(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.data, id)
	return nil
}
