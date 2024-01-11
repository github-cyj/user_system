package repository

import (
	"user_system/app/constants"
	"user_system/app/models"
	"user_system/app/request"
	"user_system/app/response"
	"user_system/utils"
)

type UserRepository struct {
}

func (repository UserRepository) GetList(params *request.UserListRequest) (userList response.UserListResponse, r *response.Response) {
	offset := (params.Page - 1) * (params.Size)
	//获取数据
	models.NewDb().Order("id desc").Offset(offset).Limit(params.Size).Find(&userList.Data).
		//获取条数
		Offset(-1).Limit(-1).Count(&userList.Total)
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
		Avatar:   params.Avatar,
		Sex:      params.Sex,
		Tel:      params.Tel,
		Email:    params.Email,
		Password: utils.MD5("123456"),
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
		Avatar:   params.Avatar,
		Sex:      params.Sex,
		Tel:      params.Tel,
		Email:    params.Email,
	}
	result := models.NewDb().Where("id = ?", id).Updates(&user)
	return result.RowsAffected, r
}

func (repository UserRepository) Delete(id uint) (deleteCount int64, r *response.Response) {
	var user models.User
	result := models.NewDb().Where("id = ?", id).Delete(&user)
	return result.RowsAffected, r
}

func (repository UserRepository) UpdatePassword(id uint, params *request.UpdatePasswordRequest) (updateCount int64, r *response.Response) {
	user, r := repository.Get(id)
	if r != nil {
		return updateCount, r
	}
	if utils.MD5(params.OldPassword) != user.Password {
		r = response.NewErrorResponseWithData(constants.ErrorBase, "原密码输入错误")
		return updateCount, r
	}
	if params.NewPassword != params.RepeatPassword {
		r = response.NewErrorResponseWithData(constants.ErrorBase, "两次密码输入不一致")
		return updateCount, r
	}
	var newPassword = utils.MD5(params.NewPassword)
	if newPassword == user.Password {
		r = response.NewErrorResponseWithData(constants.ErrorBase, "新密码不能与原密码一致")
		return updateCount, r
	}
	user.Password = newPassword
	result := models.NewDb().Where("id = ?", id).Updates(&user)
	return result.RowsAffected, r
}
