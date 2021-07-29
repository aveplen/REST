package models

type StudentResponce struct {
	StudentID           int64 `json:"studentID"`
	ScoreResponce       `json:"score"`
	SchoolResponce      `json:"school"`
	CridentialsResponce `json:"cridentials"`
}

type StudentInsert struct {
	ScoreInsert       `json:"score" binding:"required"`
	SchoolResponce    `json:"school" binding:"required"`
	CridentialsInsert `json:"cridentials" binding:"required"`
}

type StudentUpdate struct {
	ScoreResponce       `json:"score" binding:"required"`
	SchoolResponce      `json:"school" binding:"required"`
	CridentialsResponce `json:"cridentials" binding:"required"`
	StudentID           int64 `json:"studentID" binding:"required"`
}

type StudentArray struct {
	Students []*StudentResponce `json:"students"`
}

func NewStudentsArray() *StudentArray {
	return &StudentArray{
		Students: make([]*StudentResponce, 0),
	}
}
