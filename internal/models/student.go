package models

type StudentResponce struct {
	StudentID  int64                `json:"studentID"`
	Score      *ScoreResponce       `json:"score"`
	School     *SchoolResponce      `json:"school"`
	Cridetials *CridentialsResponce `json:"cridetials"`
}
