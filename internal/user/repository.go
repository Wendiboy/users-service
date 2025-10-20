package user

import (
	// "GO_API/internal/taskService"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) error
	GetAllUsers() ([]User, error)
	GetUserById(id string) (User, error)
	UpdateUser(user User) error
	DeleteUser(id string) error
	// GetUserTasks(userID string) ([]taskService.Task, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepo) GetAllUsers() ([]User, error) {

	var users []User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepo) GetUserById(id string) (User, error) {
	var user User

	err := r.db.Find(&user, "id=?", id).Error

	return user, err
}

func (r *userRepo) UpdateUser(user User) error {
	return r.db.Save(&user).Error
}

func (r *userRepo) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id=?", id).Error
}

// func (r *userRepo) GetUserTasks(userID string) ([]taskService.Task, error) {
// 	var tasks []taskService.Task

// 	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error

// 	return tasks, err
// }
