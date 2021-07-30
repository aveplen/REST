package models

type StudentResponse struct {
	StudentID           int64 `json:"studentID"`
	ScoreResponse       `json:"score"`
	SchoolResponse      `json:"school"`
	CridentialsResponse `json:"cridentials"`
}

type StudentInsert struct {
	ScoreInsert       `json:"score" binding:"required"`
	SchoolResponse    `json:"school" binding:"required"`
	CridentialsInsert `json:"cridentials" binding:"required"`
}

type StudentUpdate struct {
	ScoreResponse       `json:"score" binding:"required"`
	SchoolResponse      `json:"school" binding:"required"`
	CridentialsResponse `json:"cridentials" binding:"required"`
	StudentID           int64 `json:"studentID" binding:"required"`
}

type StudentArray struct {
	Students []*StudentResponse `json:"students"`
}

func NewStudentsArray() *StudentArray {
	return &StudentArray{
		Students: make([]*StudentResponse, 0),
	}
}
