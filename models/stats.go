package models

type Stats struct {
	Country     string `json:"Country"`
	CountryCode string `json:"CountryCode"`
	Province    string `json:"Province"`
	City        string `json:"City"`
	CityCode    string `json:"CityCode"`
	Lat         string `json:"Lat"`
	Lon         string `json:"Lon"`
	Confirmed   int    `json:"Confirmed"`
	Deaths      int    `json:"Deaths"`
	Recovered   int    `json:"Recovered"`
	Active      int    `json:"Active"`
	Date        string `json:"Date"`
}
