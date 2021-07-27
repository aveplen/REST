package models

import "time"

type CridentialsResponce struct {
	CridentialsID int64     `json:"cridentialsID"`
	FirstName     string    `json:"firstName"`
	SecondName    string    `json:"secondName"`
	Gender        string    `json:"gender"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
}

type CridentialsInsert struct {
	FirstName   string    `json:"firstName binding:"required"`
	SecondName  string    `json:"secondName binding:"required"`
	Gender      string    `json:"gender binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth binding:"required"`
}

type CridentialsUpdate struct {
	CridentialsInsert
	CridentialsID int64 `json:"cridentialsID" binding:"required"`
}
