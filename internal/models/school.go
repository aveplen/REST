package models

type School struct {
	SchoolID     int64  `json:"schoolID"`
	SchoolNumber string `json:"schoolNumber"`
	City         *City  `json:"city"`
	GeoAdress    string `json:"geoAdress"`
}
