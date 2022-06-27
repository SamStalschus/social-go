package usecase

import (
	"context"
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"
)

type DeleteUserUseCase interface {
	Execute(ctx context.Context, id int) apierrors.ApiError
}

type DeleteUser struct {
	userDao persistence.UserDao
}

func NewDeleteUsers(userDao persistence.UserDao) *DeleteUser {
	return &DeleteUser{
		userDao: userDao,
	}
}

func (deleteUser *DeleteUser) Execute(ctx context.Context, id int) apierrors.ApiError {
	err := deleteUser.userDao.DeleteById(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
