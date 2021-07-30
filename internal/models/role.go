package models

type RoleResponse struct {
	RoleID   int64  `json:"roleID"`
	RoleName string `json:"roleName"`
}

type RoleInsert struct {
	RoleName string `json:"roleName" binding:"required"`
}

type RoleUpdate struct {
	RoleInsert
	RoleID int64 `json:"roleID" binding:"required"`
}

type RoleArray struct {
	Cities []*RoleResponse `json:"roles"`
}

func NewRoleArray() *RoleArray {
	return &RoleArray{
		Cities: make([]*RoleResponse, 0),
	}
}
