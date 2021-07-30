package models

type SchoolResponse struct {
	SchoolID     int64  `json:"schoolID"`
	SchoolNumber string `json:"schoolNumber"`
	CityResponse `json:"city"`
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
	Schools []*SchoolResponse `json:"schools"`
}

func NewSchoolArray() *SchoolArray {
	return &SchoolArray{
		Schools: make([]*SchoolResponse, 0),
	}
}
