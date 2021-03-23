package trips

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/enums"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/posts"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/users"
)

type MultiLineString struct {
	Type        string        `json:"type",default:"MultiLineString"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type TripDomains struct {
	Id           string              `json:"id"`
	StartDate    string              `json:"startDate"`
	EndDate      string              `json:"endDate"`
	PublicStatus enums.TripViewType  `json:"publicStatus"`
	CompanyName  string              `json:"companyName,omitempty"`
	ShipNumber   string              `json:"shipNumber,omitempty"`
	Publisher    users.UserDomains   `json:"publisher"`
	Viewers      []users.UserDomains `json:"viewers,omitempty"`
	Posts        []posts.PostDomains `json:"posts,omitempty"`
	CreatedAt    string              `json:"createdAt"`
	UpdatedAt    string              `json:"updatedAt"`
}
