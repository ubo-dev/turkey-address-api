package model

type City struct {
	Id       int    `json:"cityId"`
	CityName string `json:"cityName"`
}

type District struct {
	Id           int    `json:"districtId"`
	DistrictName string `json:"districtName"`
	CityId       int    `json:"cityId"`
}

type Neighbourhood struct {
	Id                int    `json:"neighbourhoodId"`
	NeighbourhoodName string `json:"neighbourhoodName"`
	DistrictId        int    `json:"districtId"`
}
