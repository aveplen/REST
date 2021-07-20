package models

type School struct {
	SchoolID     int64  `json:"schoolID"`
	SchoolNumber string `json:"schoolNumber"`
	CityID       int64  `json:"dityID"`
	GeoAdress    string `json:"geoAdress"`
}
