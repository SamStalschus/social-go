package persistence

import (
	"log"
	"social-go/cmd/api/core/model"
	"time"

	"github.com/jinzhu/gorm"
)

type user struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (userDao *UserDao) Save(userToSave *model.User) (*model.User, error) {

	newUser := &user{
		CreatedAt: time.Now(),
		Name:      userToSave.Name,
		Username:  userToSave.Username,
		Password:  userToSave.Password,
		Email:     userToSave.Email,
	}

	if err := userDao.db.Create(newUser).Error; err != nil {
		log.Fatal("Error to create a new user.")
	}

	return parseEntity(newUser), nil
}

func parseEntity(row *user) *model.User {
	return &model.User{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		Name:      row.Name,
		Username:  row.Username,
		Password:  row.Password,
		Email:     row.Email,
	}
}
