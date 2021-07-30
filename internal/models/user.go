package models

type UserResponse struct {
	UserID            int64            `json:"userID"`
	Email             string           `json:"email"`
	EncryptedPassword string           `json:"encryptedPassword"`
	Role              RoleResponse     `json:"role"`
	Student           *StudentResponse `json:"student,omitempty"`
}

type UserInsert struct {
	Email             string           `json:"email"`
	EncryptedPassword string           `json:"encryptedPassword"`
	Role              RoleResponse     `json:"role"`
	Student           *StudentResponse `json:"student,omitempty"`
}

type UserUpdate struct {
	UserInsert
	UserID int64 `json:"userID" binding:"required"`
}

type UserArray struct {
	Users []*UserResponse `json:"users"`
}

func NewUserArray() *UserArray {
	return &UserArray{
		Users: make([]*UserResponse, 0),
	}
}
