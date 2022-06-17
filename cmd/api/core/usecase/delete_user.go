package usecase

import (
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"
)

type DeleteUserUseCase interface {
	Execute(user string)
}

type DeleteUser struct {
	userDao persistence.UserDao
}

func NewDeleteUsers(userDao persistence.UserDao) *DeleteUser {
	return &DeleteUser{
		userDao: userDao,
	}
}

func (deleteUser *DeleteUser) Execute(id int) apierrors.ApiError {
	err := deleteUser.userDao.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}
