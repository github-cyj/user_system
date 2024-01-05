package response

type AuthLoginResponse struct {
	Token string `json:"token"`
}

type UserInfoResponse struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}
