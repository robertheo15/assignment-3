package structs

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type DataJSON struct {
	Status `json:"status"`
}

type Weather struct {
	WaterValue  int    `json:"water_value"`
	WaterStatus string `json:"water_status"`
	WindValue   int    `json:"wind_value"`
	WindStatus  string `json:"wind_status"`
}
