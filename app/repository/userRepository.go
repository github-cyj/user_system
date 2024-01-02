package repository

import (
	"user_system/app/constants"
	"user_system/app/models"
	"user_system/app/request"
	"user_system/app/response"
)

type UserRepository struct {
}

func (repository UserRepository) GetList(params *request.UserListRequest) (userList []models.User, r *response.Response) {
	offset := (params.Page - 1) * (params.Size)
	models.NewDb().Offset(offset).Limit(params.Size).Find(&userList)
	return userList, r
}

func (repository UserRepository) Get(id uint) (user models.User, r *response.Response) {
	models.NewDb().Where("id = ?", id).First(&user)
	if user.ID == 0 {
		r = response.NewErrorResponseWithData(constants.ErrorNotExits, "用户")
	}
	return user, r
}

func (repository UserRepository) Add(params *request.UserAddRequest) (id uint, r *response.Response) {
	user := models.User{
		Username: params.Username,
	}
	models.NewDb().Create(&user)
	return user.ID, r
}

func (repository UserRepository) Edit(id uint, params *request.UserEditRequest) (updateCount int64, r *response.Response) {
	_, r = repository.Get(id)
	if r != nil {
		return updateCount, r
	}
	user := models.User{
		Username: params.Username,
	}
	result := models.NewDb().Where("id = ?", id).Updates(&user)
	return result.RowsAffected, r
}

func (repository UserRepository) Delete(id uint) (deleteCount int64, r *response.Response) {
	var user models.User
	result := models.NewDb().Where("id = ?", id).Delete(&user)
	return result.RowsAffected, r
}
