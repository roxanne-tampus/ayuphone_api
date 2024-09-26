package models

type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Province struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RegionID int    `json:"region_id"`
}

type Municipality struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ProvinceID int    `json:"province_id"`
}

type Barangay struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	MunicipalityID int    `json:"municipality_id"`
}
