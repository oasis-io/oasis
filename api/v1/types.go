package v1

import "oasis/db/model"

type PageInfo struct {
	CurrentPage int `json:"currentPage" form:"currentPage"` // 页码
	PageSize    int `json:"pageSize" form:"pageSize"`       // 每页大小
}

type PageResponse struct {
	Data        interface{} `json:"data"`
	Total       int64       `json:"total"`
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
}

type UserResponse struct {
	User model.User `json:"user"`
}

type UserRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RoleResponse struct {
	Name string `json:"name"`
}
