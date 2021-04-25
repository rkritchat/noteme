package repository

import (
	"github.com/gobuffalo/pop"
	"time"
)

type UserEntity struct {
	ID        int64      `db:"id"`
	Username  string     `db:"username"`
	Firstname string     `db:"firstname"`
	Lastname  string     `db:"lastname"`
	Email     string     `db:"email"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type UserRepository interface {
	Create(entity UserEntity) error
	FindByUsername(username string) (*UserEntity, error)
}

type userRepository struct {
	DB *pop.Connection
}

func NewUserRepository(DB *pop.Connection) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (repo userRepository) FindByUsername(username string) (*UserEntity, error) {
	var r UserEntity
	return &r, repo.DB.Where("username = ?", username).First(&r)
}

func (repo userRepository) Create(entity UserEntity) error {
	return repo.DB.Create(&entity)
}

func (UserEntity) TableName() string {
	return "user"
}
