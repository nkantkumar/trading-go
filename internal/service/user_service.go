package service

import (
	"trading-go/internal/model"
	"trading-go/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserDetails(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) RegisterUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) RemoveUser(id int) error {
	return s.repo.DeleteUser(id)
}
