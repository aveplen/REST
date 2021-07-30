package models

type GenderResponse struct {
	GenderID   int64  `json:"genderID"`
	GenderName string `json:"genderName"`
}

type GenderInsert struct {
	GenderName string `json:"genderName" binding:"required"`
}

type GenderUpdate struct {
	GenderInsert
	GenderID int64 `json:"genderID" binding:"required"`
}

type GenderArray struct {
	Cities []*GenderResponse `json:"genders"`
}

func NewGenderArray() *GenderArray {
	return &GenderArray{
		Cities: make([]*GenderResponse, 0),
	}
}
