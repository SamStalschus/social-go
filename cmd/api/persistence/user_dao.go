package persistence

import (
	"errors"
	"log"
	"social-go/cmd/api/core/model"
	apierrors "social-go/cmd/api/utils/err"
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
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (userDao *UserDao) Save(userToSave *model.User) (*model.User, apierrors.ApiError) {

	newUser := &user{
		CreatedAt: time.Now(),
		Name:      userToSave.Name,
		Username:  userToSave.Username,
		Password:  userToSave.Password,
		Email:     userToSave.Email,
	}

	if err := userDao.db.Create(newUser).Error; err != nil {
		return nil, apierrors.NewBadRequestApiError("Error to create a new user." + err.Error())
	}

	return parseEntity(newUser), nil
}

func (userDao *UserDao) Update(userToUpdate *model.User) (*model.User, error) {
	var user user

	// var dbError APIERROR

	err := userDao.db.First(&user, userToUpdate.ID).Error
	switch {
	case gorm.IsRecordNotFoundError(err):
		// TODO: Error handling
		// DBERROR
		log.Fatal("User not found!")
	case err != nil:
		// TODO: Error handling
		// DBERROR
		log.Fatal("Error to finding user!")
	default:
		user.Name = userToUpdate.Name
		user.Username = userToUpdate.Username
		if err = userDao.db.Save(&user).Error; err != nil {
			// TODO: Error handling
			log.Fatal("Error to finding user!")
		}
	}

	// if dbError != nil {
	// 	return nil, dbError
	// }

	return parseEntity(&user), nil
}

func (userDao *UserDao) Get(id int) (*model.User, error) {
	var user user

	var dbError error

	if err := userDao.db.First(&user, id).Error; gorm.IsRecordNotFoundError(err) {
		dbError = errors.New("Error handling")
	} else if err != nil {
		dbError = errors.New("Error handling")
	}

	if dbError != nil {
		return nil, dbError
	}

	return parseEntity(&user), nil

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
