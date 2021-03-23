package posts

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/enums"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/trips"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/users"
)

type PostDomains struct {
	Id           string              `json:"id"`
	Content      string              `json:"content"`
	Images       []string            `json:"images,omitempty"`
	PublicStatus enums.TripViewType  `json:"publicStatus"`
	Publisher    users.UserDomains   `json:"publisher"`
	LikeUsers    []users.UserDomains `json:"likeUsers,omitempty"`
	Trip         trips.TripDomains   `json:"trip"`
	CreatedAt    string              `json:"createdAt"`
	UpdatedAt    string              `json:"updatedAt"`
}
