package models

type Student struct {
	StudentID  int64        `json:"studentID"`
	Score      *Score       `json:"score"`
	School     *School      `json:"school"`
	Cridetials *Cridentials `json:"cridetials"`
}
