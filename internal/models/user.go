package models

type (
	UserResponse struct {
		UserID            int64            `json:"user_id"`
		Email             string           `json:"email"`
		EncryptedPassword string           `json:"encrypted_password"`
		Role              string           `json:"role"`
		StudResponse      *StudentResponse `json:"student,omitempty"`
	}

	UserInsert struct {
		Email             string `json:"email" binding:"required"`
		EncryptedPassword string `json:"encrypted_password" binding:"required"`
	}

	UserUpdate struct {
		UserID int64 `json:"user_id" binding:"required"`
		UserInsert
	}

	UserAttach struct {
		UserID    int64
		StudentID int64
	}

	UserRole struct {
		UserID int64  `json:"user_id"`
		Role   string `json:"role"`
	}

	UserAuth struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	UserExistance struct {
		Email string
	}
)
