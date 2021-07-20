package models

import "time"

type Cridentials struct {
	CridentialsID int64     `json:"cridentialsID"`
	FirstName     string    `json:"firstName"`
	SecondName    string    `json:"secondName"`
	Gender        string    `json:"gender"`
	DateOfBirth   time.Time `json:"dateOfBirth"`

	// remove omitempty for authentication
	Email string `json:"email,omitempty"`

	// add salted password for authentication
}
