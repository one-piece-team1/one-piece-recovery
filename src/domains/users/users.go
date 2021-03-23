package users

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/enums"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/posts"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/trips"
)

type UserDomains struct {
	Id             string               `json:"id"`
	Role           enums.UserGenderType `json:"role"`
	ExpiredDate    string               `json:"expiredDate"`
	DiamondCoin    int                  `json:"diamondCoin"`
	GoldCoin       int                  `json:"goldCoin"`
	Username       string               `json:"username"`
	Email          string               `json:"email"`
	Password       string               `json:"password"`
	Salt           string               `json:"salt"`
	Status         bool                 `json:"status"`
	Gender         enums.UserGenderType `json:"gender,omitempty"`
	Age            int                  `json:"age,omitempty"`
	Desc           string               `json:"desc,omitempty"`
	ProfileImage   string               `json:"profileImage,omitempty"`
	Trips          []trips.TripDomains  `json:"trips,omitempty"`
	Views          []trips.TripDomains  `json:"views,omitempty"`
	Posts          []posts.PostDomains  `json:"posts,omitempty"`
	LikePosts      []posts.PostDomains  `json:"likePosts,omitempty"`
	Followers      []UserDomains        `json:"followers,omitempty"`
	Followings     []UserDomains        `json:"followings,omitempty"`
	BlockLists     []UserDomains        `json:"blockLists,omitempty"`
	FollowerCount  int                  `json:"followerCount"`
	FollowingCount int                  `json:"followingCount"`
}
