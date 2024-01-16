package controllers

import (
	"resq-be/model"
	"resq-be/usecases"
	"resq-be/utils"

	"github.com/gin-gonic/gin"
)

type user struct {
	userUsecase usecases.User
}

func NewUser(userUsecase usecases.User) *user {
	return &user{userUsecase}
}

func (u *user) Register(ctx *gin.Context) {
	var arg model.UserRegister
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		utils.SetError(ctx, err, utils.ErrUnprocessableEntity)
		return
	}
	message, err := u.userUsecase.Register(ctx, &arg)
	if err != nil {
		utils.SetError(ctx, err, message)
		return
	}
	utils.Success(ctx, 201, gin.H{"status": "success"})
}

func (u *user) Login(ctx *gin.Context) {
	var arg model.UserLogin
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		utils.SetError(ctx, err, utils.ErrUnprocessableEntity)
		return
	}
	message, err := u.userUsecase.Login(ctx, &arg)
	if err != nil {
		utils.SetError(ctx, err, message)
		return
	}
	utils.Success(ctx, 200, gin.H{
		"status": "OTP sent to you number",
	})
}

func (u *user) Update(ctx *gin.Context) {
	var arg model.UserUpdate
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		utils.SetError(ctx, err, utils.ErrUnprocessableEntity)
		return
	}
	id := ctx.Param("id")
	message, err := u.userUsecase.Update(ctx, id, &arg)
	if err != nil {
		utils.SetError(ctx, err, message)
		return
	}
	utils.Success(ctx, 200, nil)
}

func (u *user) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	message, err := u.userUsecase.Delete(ctx, id)
	if err != nil {
		utils.SetError(ctx, err, message)
		return
	}
	utils.Success(ctx, 200, nil)
}

func (u *user) Profile(ctx *gin.Context) {
	id := ctx.Param("id")
	user, message, err := u.userUsecase.Profile(ctx, id)
	if err != nil {
		utils.SetError(ctx, err, message)
		return
	}
	utils.Success(ctx, 200, user)
}
