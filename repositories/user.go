package repositories

import (
	"context"
	"resq-be/model"

	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*model.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	FindAllAdmin(ctx context.Context) ([]*model.User, error)
	FindAllNotAdmin(ctx context.Context) ([]*model.User, error)
}

func NewUser(DB *gorm.DB) UserRepository {
	return &userRepo{DB}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	tx := r.DB.Begin()
	if err := tx.Create(user).WithContext(ctx).Error; err != nil {
		return tx.Rollback().Error
	}
	return tx.Commit().Error
}

func (r *userRepo) Update(ctx context.Context, user *model.User) error {
	tx := r.DB.Begin()
	if err := tx.Save(user).WithContext(ctx).Error; err != nil {
		return tx.Rollback().Error
	}
	return tx.Commit().Error
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&model.User{}, id).WithContext(ctx).Error; err != nil {
		return tx.Rollback().Error
	}
	return tx.Commit().Error
}

func (r *userRepo) FindByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) FindAllAdmin(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.WithContext(ctx).Where("is_admin = ?", true).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) FindAllNotAdmin(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.WithContext(ctx).Where("is_admin = ?", false).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
