package locations

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/countries"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/enums"
)

type GeoPoint struct {
	Type        string    `json:"type", default:"Point"`
	Coordinates []float64 `json:"coordinates"`
}

type LocationDomains struct {
	Id           string                   `json:"id"`
	Point        GeoPoint                 `json:"point"`
	PointSrid    GeoPoint                 `json:"pointSrid"`
	Lat          float64                  `json:"lat"`
	Lon          float64                  `json:"lon"`
	Type         enums.LocationType       `json:"type"`
	LocationName string                   `json:"locationName"`
	Country      countries.CountryDomains `json:"country"`
	CreatedAt    string                   `json:"createdAt"`
	UpdatedAt    string                   `json:"updatedAt"`
}
