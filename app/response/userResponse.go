package response

import "user_system/app/models"

type UserListResponse struct {
	Total int64         `json:"total" `
	Data  []models.User `json:"data"`
}
