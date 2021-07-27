package models

type SchoolResponce struct {
	SchoolID     int64         `json:"schoolID"`
	SchoolNumber string        `json:"schoolNumber"`
	City         *CityResponce `json:"city"`
	GeoAdress    string        `json:"geoAdress"`
}

type SchoolInsert struct {
	SchoolNumber string `json:"schoolNumber binding:"required"`
	CityName     string `json:"cityName binding:"required"`
	GeoAdress    string `json:"geoAdress binding:"required"`
}

type SchoolUpdate struct {
	SchoolInsert
	SchoolID int64 `json:"schoolID" binding:"required"`
}
