package models

import "time"

type CridentialsResponse struct {
	CridentialsID int64     `json:"cridentialsID"`
	FirstName     string    `json:"firstName"`
	SecondName    string    `json:"secondName"`
	Gender        string    `json:"gender"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
}

type CridentialsInsert struct {
	FirstName   string    `json:"firstName" binding:"required"`
	SecondName  string    `json:"secondName" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
}

type CridentialsUpdate struct {
	CridentialsInsert
	CridentialsID int64 `json:"cridentialsID" binding:"required"`
}

type CridentialsArray struct {
	CridentialsArr []*CridentialsResponse `json:"cridentialsArr"`
}

func NewCridentialsArray() *CridentialsArray {
	return &CridentialsArray{
		CridentialsArr: make([]*CridentialsResponse, 0),
	}
}
