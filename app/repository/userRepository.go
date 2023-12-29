package repository

import (
	"user_system/app/models"
	"user_system/app/request"
	"user_system/app/response"
)

type UserRepository struct {
}

func (repository UserRepository) GetList(params *request.UserListRequest) []models.User {
	offset := (params.Page - 1) * (params.Size)
	var userList []models.User
	models.NewDb().Offset(offset).Limit(params.Size).Find(&userList)
	return userList
}

func (repository UserRepository) Get(id uint) models.User {
	var user models.User

	models.NewDb().Where("id = ?", id).First(&user)

	return user
}

func (repository UserRepository) Add(params *request.UserAddRequest) uint {
	user := models.User{
		Username: params.Username,
	}
	models.NewDb().Create(&user)
	return user.ID
}

func (repository UserRepository) Edit(id uint, params *request.UserEditRequest) int64 {
	user := models.User{
		Username: params.Username,
	}
	result := models.NewDb().Where("id = ?", id).Updates(&user)
	return result.RowsAffected
}

func (repository UserRepository) Delete(id uint) response.Response {
	var user models.User
	result := models.NewDb().Where("id = ?", id).Delete(&user)

	return response.NewSuccess(result.RowsAffected)
}
