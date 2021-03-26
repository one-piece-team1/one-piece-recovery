package recoveries

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/enums"
)

type ChatSendStatus = string

const (
	FAIL    ChatSendStatus = "fail"
	SENDING                = "sending"
	FINISH                 = "finish"
)

type ChatStatus = string

const (
	READ   ChatStatus = "read"
	UNREAD            = "unread"
)

type ChatRoomType = string

// Conflict with type type move to here
const (
	PUBLIC  ChatRoomType = "public"
	PRIVATE              = "private"
	GROUP                = "group"
)

type MultiLineString struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type GeoPoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type ChatParitcipateDomains struct {
	Id        string           `json:"id"`
	ChatRoom  *ChatRoomDomains `json:"chatroom,omitempty"`
	Chats     *[]ChatDomains   `json:"chats,omitempty"`
	Users     *[]UserDomains   `json:"users,omitempty"`
	CreatedAt string           `json:"createdAt"`
	UpdatedAt string           `json:"updatedAt"`
}

type ChatRoomDomains struct {
	Id              string                  `json:"id"`
	Name            string                  `json:"name"`
	Type            ChatRoomType            `json:"type"`
	ChatParticipate *ChatParitcipateDomains `json:"chatParticipate,omitempty"`
	CreatedAt       string                  `json:"createdAt"`
	UpdatedAt       string                  `json:"updatedAt"`
}

type ChatDomains struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"name"`
	SendStatus      ChatSendStatus          `json:"sendStatus"`
	ChatParticipate *ChatParitcipateDomains `json:"chatParticipate,omitempty"`
	CreatedAt       string                  `json:"createdAt"`
	UpdatedAt       string                  `json:"updatedAt"`
}

type CountryDomains struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type LocationDomains struct {
	Id           string             `json:"id"`
	Point        GeoPoint           `json:"point"`
	PointSrid    GeoPoint           `json:"pointSrid"`
	Lat          float64            `json:"lat"`
	Lon          float64            `json:"lon"`
	Type         enums.LocationType `json:"type"`
	LocationName string             `json:"locationName"`
	Country      CountryDomains     `json:"country"`
	CreatedAt    string             `json:"createdAt"`
	UpdatedAt    string             `json:"updatedAt"`
}

type PostDomains struct {
	Id           string             `json:"id"`
	Content      string             `json:"content"`
	Images       []string           `json:"images,omitempty"`
	PublicStatus enums.TripViewType `json:"publicStatus"`
	Publisher    UserDomains        `json:"publisher"`
	LikeUsers    []UserDomains      `json:"likeUsers,omitempty"`
	Trip         TripDomains        `json:"trip"`
	CreatedAt    string             `json:"createdAt"`
	UpdatedAt    string             `json:"updatedAt"`
}

type TripDomains struct {
	Id           string             `json:"id"`
	StartDate    string             `json:"startDate"`
	EndDate      string             `json:"endDate"`
	PublicStatus enums.TripViewType `json:"publicStatus"`
	CompanyName  string             `json:"companyName,omitempty"`
	ShipNumber   string             `json:"shipNumber,omitempty"`
	Publisher    UserDomains        `json:"publisher"`
	Viewers      []UserDomains      `json:"viewers,omitempty"`
	Posts        []PostDomains      `json:"posts,omitempty"`
	CreatedAt    string             `json:"createdAt"`
	UpdatedAt    string             `json:"updatedAt"`
}

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
	Trips          []TripDomains        `json:"trips,omitempty"`
	Views          []TripDomains        `json:"views,omitempty"`
	Posts          []PostDomains        `json:"posts,omitempty"`
	LikePosts      []PostDomains        `json:"likePosts,omitempty"`
	Followers      []UserDomains        `json:"followers,omitempty"`
	Followings     []UserDomains        `json:"followings,omitempty"`
	BlockLists     []UserDomains        `json:"blockLists,omitempty"`
	FollowerCount  int                  `json:"followerCount"`
	FollowingCount int                  `json:"followingCount"`
}
