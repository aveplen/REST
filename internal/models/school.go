package models

type SchoolResponce struct {
	SchoolID     int64  `json:"schoolID"`
	SchoolNumber string `json:"schoolNumber"`
	CityID       int64  `json:"cityID"`
	GeoAddress   string `json:"geoAddress"`
}

type SchoolInsert struct {
	SchoolNumber string `json:"schoolNumber" binding:"required"`
	CityName     string `json:"cityName" binding:"required"`
	GeoAddress   string `json:"geoAddress" binding:"required"`
}

type SchoolUpdate struct {
	SchoolInsert
	SchoolID int64 `json:"schoolID" binding:"required"`
}

type SchoolArray struct {
	Schools []*SchoolResponce `json:"schools"`
}

func NewSchoolArray() *SchoolArray {
	return &SchoolArray{
		Schools: make([]*SchoolResponce, 0),
	}
}
