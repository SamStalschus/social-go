package persistence

import (
	"context"
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

func (userDao *UserDao) Save(ctx context.Context, userToSave *model.User) (*model.User, apierrors.ApiError) {

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

func (userDao *UserDao) Update(ctx context.Context, userToUpdate *model.User) (*model.User, apierrors.ApiError) {
	var user user

	var dbError apierrors.ApiError

	err := userDao.db.First(&user, userToUpdate.ID).Error
	switch {
	case gorm.IsRecordNotFoundError(err):
		dbError = apierrors.NewNotFoundApiError("User not found!")
	case err != nil:
		dbError = apierrors.NewBadRequestApiError("Error to finding user!")
	default:
		user.Name = userToUpdate.Name
		user.Username = userToUpdate.Username
		if err = userDao.db.Save(&user).Error; err != nil {
			dbError = apierrors.NewBadRequestApiError("Error to update user!")
		}
	}

	if dbError != nil {
		return nil, dbError
	}

	return parseEntity(&user), nil
}

func (userDao *UserDao) Get(ctx context.Context, id int) (*model.User, apierrors.ApiError) {
	var user user
	var dbError apierrors.ApiError

	if err := userDao.db.First(&user, id).Error; gorm.IsRecordNotFoundError(err) {
		dbError = apierrors.NewNotFoundApiError("User not found!")
	} else if err != nil {
		dbError = apierrors.NewBadRequestApiError("Error to finding user!")
	}

	if dbError != nil {
		return nil, dbError
	}

	return parseEntity(&user), nil
}

func (userDao *UserDao) GetByUsernameWithPassword(ctx context.Context, username string) (*model.User, apierrors.ApiError) {
	var user user
	var dbError apierrors.ApiError

	if err := userDao.db.Where("username = ?", username).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		dbError = apierrors.NewNotFoundApiError("User not found!")
	} else if err != nil {
		dbError = apierrors.NewBadRequestApiError("Error to finding user!")
	}

	if dbError != nil {
		return nil, dbError
	}

	return &model.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
	}, nil
}

func (userDao *UserDao) GetAll(ctx context.Context) (*[]model.User, apierrors.ApiError) {
	var usersDao []user
	var users []model.User

	if err := userDao.db.Find(&usersDao).Error; err != nil {
		return nil, apierrors.NewInternalServerError()
	}

	for _, user := range usersDao {
		users = append(users, *parseEntity(&user))
	}

	return &users, nil
}

func (userDao *UserDao) DeleteById(ctx context.Context, id int) apierrors.ApiError {
	var user user

	err := userDao.db.First(&user, id).Error
	switch {
	case gorm.IsRecordNotFoundError(err):
		return apierrors.NewNotFoundApiError("User not found!")
	case err != nil:
		return apierrors.NewBadRequestApiError("Error to delete user!")
	default:
		if err = userDao.db.Delete(&user, id).Error; err != nil {
			return apierrors.NewBadRequestApiError("Error to delete user!")
		}
	}

	return nil
}

func parseEntity(row *user) *model.User {
	return &model.User{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		Name:      row.Name,
		Username:  row.Username,
		Password:  "",
		Email:     row.Email,
	}
}
