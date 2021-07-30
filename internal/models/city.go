package models

type CityResponse struct {
	CityID   int64  `json:"cityId"`
	CityName string `json:"cityName"`
}

type CityInsert struct {
	CityName string `json:"cityName" binding:"required"`
}

type CityUpdate struct {
	CityInsert
	CityID int64 `json:"cityId" binding:"required"`
}

type CityArray struct {
	Cities []*CityResponse `json:"cities"`
}

func NewCityArray() *CityArray {
	return &CityArray{
		Cities: make([]*CityResponse, 0),
	}
}
