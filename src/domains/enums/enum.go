package enums

type UserRoleType = string

const (
	TRAIL UserRoleType = "trail"
	USER               = "user"
	VIP1               = "vip1"
	VIP2               = "vip2"
	ADMIN              = "admin"
)

type UserGenderType = string

const (
	MALE   UserGenderType = "male"
	FEMALE                = "female"
)

type TripViewType = string

const (
	PUBLIC TripViewType = "public"
	FRIEND              = "friend"
	SELF                = "self"
)

type LocationType = string

const (
	COUNTRY LocationType = "country"
	CITY                 = "city"
	SCENE                = "scene"
	PORT                 = "port"
	TURN                 = "turn"
)

type LocationQueryType = string

const (
	SPECIFIC LocationQueryType = "specific"
	RANGE                      = "range"
)
