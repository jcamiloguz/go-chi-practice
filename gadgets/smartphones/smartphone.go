package smartphones

//Smartphone model structure  for smartphone
type Smartphone struct {
	ID            int64
	Name          string
	Price         int
	CountryOrigin string
	Os            string
}
type CreateSmartphone struct {
	Name          string `json:"name"`
	Price         int    `json:"Price"`
	CountryOrigin string `json:"country_origin"`
	Os            string `json:"os"`
}
