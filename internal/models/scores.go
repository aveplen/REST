package models

type Score struct {
	ScoreID          int64 `json:"scoresID"`
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
