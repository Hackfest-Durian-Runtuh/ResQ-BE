package usecases

import (
	"context"
	"resq-be/model"
	"resq-be/repositories"
	"resq-be/utils"
	"strconv"
)

type user struct {
	userRepo repositories.UserRepository
}

type User interface {
	Register(ctx context.Context, arg *model.UserRegister) (string, error)
	Login(ctx context.Context, arg *model.UserLogin) (string, error)
	Update(ctx context.Context, id string, arg *model.UserUpdate) (string, error)
	Delete(ctx context.Context, id string) (string, error)
	Profile(ctx context.Context, id string) (*model.User, string, error)
}

func NewUser(userRepo repositories.UserRepository) User {
	return &user{userRepo}
}

func (u *user) Register(ctx context.Context, arg *model.UserRegister) (string, error) {
	if err := arg.Validate(); err != nil {
		return utils.ErrUnprocessableEntity, err
	}
	if err := u.userRepo.Create(ctx, arg.ToUser()); err != nil {
		return utils.ErrInternalServerError, err
	}
	return "", nil
}

func (u *user) Login(ctx context.Context, arg *model.UserLogin) (string, error) {
	user, err := u.userRepo.FindByPhoneNumber(ctx, arg.PhoneNumber)
	if err != nil {
		return utils.ErrInternalServerError, err
	}
	if user == nil {
		return utils.ErrNotFound, err
	}
	return "", nil
}

func (u *user) Update(ctx context.Context, id string, arg *model.UserUpdate) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrBadRequest, err
	}
	user, err := u.userRepo.FindByID(ctx, idInt)
	if err != nil {
		return utils.ErrInternalServerError, err
	}
	if user == nil {
		return utils.ErrNotFound, err
	}
	if arg.Name != "" {
		user.Name = arg.Name
	}
	if arg.NIK != "" {
		user.NIK = arg.NIK
	}
	if arg.PhoneNumber != "" {
		user.PhoneNumber = arg.PhoneNumber
	}
	if arg.Image != nil {
		user.Image = arg.ToUser().Image
	}
	if err := u.userRepo.Update(ctx, arg.ToUser()); err != nil {
		return utils.ErrInternalServerError, err
	}
	return "", nil
}

func (u *user) Delete(ctx context.Context, id string) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrBadRequest, err
	}
	user, err := u.userRepo.FindByID(ctx, idInt)
	if err != nil {
		return utils.ErrInternalServerError, err
	}
	if user == nil {
		return utils.ErrNotFound, err
	}
	if err := u.userRepo.Delete(ctx, idInt); err != nil {
		return utils.ErrInternalServerError, err
	}
	return "", nil
}

func (u *user) Profile(ctx context.Context, id string) (*model.User, string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, utils.ErrBadRequest, err
	}
	user, err := u.userRepo.FindByID(ctx, idInt)
	if err != nil {
		return nil, utils.ErrInternalServerError, err
	}
	if user == nil {
		return nil, utils.ErrNotFound, err
	}
	return user, "", nil
}
