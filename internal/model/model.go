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

type Neighboorhood struct {
	Id                int    `json:"neighboorhoodId"`
	NeighboorhoodName string `json:"neighboorhoodName"`
	DistrictId        int    `json:"districtId"`
}
