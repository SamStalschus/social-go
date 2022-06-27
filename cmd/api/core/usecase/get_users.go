package usecase

import (
	"context"
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"
)

type GetUsersUseCase interface {
	Execute(ctx context.Context) (*[]model.User, apierrors.ApiError)
}

type GetUsers struct {
	userDao persistence.UserDao
}

func NewGetUsers(userDao persistence.UserDao) *GetUsers {
	return &GetUsers{
		userDao: userDao,
	}
}

func (getUsers *GetUsers) Execute(ctx context.Context) (*[]model.User, apierrors.ApiError) {
	users, err := getUsers.userDao.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
