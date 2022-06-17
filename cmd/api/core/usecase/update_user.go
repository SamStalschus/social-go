package usecase

import (
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"
)

type UpdateUserUseCase interface {
	Execute(user string)
}

type UpdateUser struct {
	userDao persistence.UserDao
}

func NewUpdateUser(userDao persistence.UserDao) *UpdateUser {
	return &UpdateUser{
		userDao: userDao,
	}
}

func (updateUser *UpdateUser) Execute(user *model.User) (*model.User, apierrors.ApiError) {
	userUpdated, err := updateUser.userDao.Update(user)

	if err != nil {
		return nil, err
	}

	return userUpdated, nil
}
