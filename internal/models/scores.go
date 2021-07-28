package models

type ScoreResponce struct {
	ScoreID          int64 `json:"scoreID"`
	Mathematics      *int  `json:"mathematics,omitempty"`
	Russian          *int  `json:"russian,omitempty"`
	Physics          *int  `json:"physics,omitempty"`
	ComputerScience  *int  `json:"computerScience,omitempty"`
	Literature       *int  `json:"literature,omitempty"`
	SocialScience    *int  `json:"socialScience,omitempty"`
	History          *int  `json:"history,omitempty"`
	Biology          *int  `json:"biology,omitempty"`
	GeographyScience *int  `json:"geographyScience,omitempty"`
}

type ScoreInsert struct {
	Mathematics      *int `json:"mathematics,omitempty"`
	Russian          *int `json:"russian,omitempty"`
	Physics          *int `json:"physics,omitempty"`
	ComputerScience  *int `json:"computerScience,omitempty"`
	Literature       *int `json:"literature,omitempty"`
	SocialScience    *int `json:"socialScience,omitempty"`
	History          *int `json:"history,omitempty"`
	Biology          *int `json:"biology,omitempty"`
	GeographyScience *int `json:"geographyScience,omitempty"`
}

type ScoreUpdate struct {
	ScoreInsert
	ScoreID int64 `json:"scoreID" binding:"required"`
}

type ScoreArray struct {
	Scores []*ScoreResponce `json:"scores"`
}

func NewScoreArray() *ScoreArray {
	return &ScoreArray{
		Scores: make([]*ScoreResponce, 0),
	}
}
