package models

type StudentResponce struct {
	StudentID  int64                `json:"studentID"`
	Score      *ScoreResponce       `json:"score"`
	School     *SchoolResponce      `json:"school"`
	Cridetials *CridentialsResponce `json:"cridetials"`
}

type StudentInsert struct {
	Score      *ScoreResponce       `json:"score" binding:"required"`
	School     *SchoolResponce      `json:"school" binding:"required"`
	Cridetials *CridentialsResponce `json:"cridetials" binding:"required"`
}

type StudentUpdate struct {
	StudentInsert
	StudentID int64 `json:"studentID" binding:"required"`
}

type StudentArray struct {
	Students []*StudentResponce `json:"students"`
}

func NewStudentsArray() *StudentArray {
	return &StudentArray{
		Students: make([]*StudentResponce, 0),
	}
}
