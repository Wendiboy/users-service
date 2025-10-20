package user

import "math/rand"

type Service interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserById(id uint32) (User, error)
	UpdateUser(updatedUser User) (User, error)
	DeleteUser(id uint32) error
}

type userService struct {
	repo UserRepository
}

func NewService(r UserRepository) Service {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user User) (User, error) {

	user.Id = rand.Uint32()
	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserById(id uint32) (User, error) {
	return s.repo.GetUserById(id)
}

func (s *userService) UpdateUser(updatedUser User) (User, error) {
	user, err := s.repo.GetUserById(updatedUser.Id)
	if err != nil {
		return User{}, err
	}

	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint32) error {
	return s.repo.DeleteUser(id)
}
