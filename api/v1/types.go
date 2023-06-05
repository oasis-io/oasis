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

type UserRequest struct {
	model.User
	Roles []string `json:"roles"`
}

type UserRes struct {
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Phone    string         `json:"phone"`
	Password string         `json:"password"`
	Roles    []RoleResponse `json:"roles"`
}

type MenuRequest struct {
	Name  string       `json:"name"`
	Apis  []model.Api  `json:"apis"`
	Menus []model.Menu `json:"menus"`
}

type RoleRequest struct {
	Name string `json:"name"`
}

type RoleResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
