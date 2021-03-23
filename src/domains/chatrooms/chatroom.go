package chatrooms

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/chatparticipates"
)

type ChatRoomType = string

// Conflict with type type move to here
const (
	PUBLIC  ChatRoomType = "public"
	PRIVATE              = "private"
	GROUP                = "group"
)

type ChatRoomDomains struct {
	Id              string                                  `json:"id"`
	Name            string                                  `json:"name"`
	Type            ChatRoomType                            `json:"type"`
	ChatParticipate chatparticipates.ChatParitcipateDomains `json:"chatParticipate"`
	CreatedAt       string                                  `json:"createdAt"`
	UpdatedAt       string                                  `json:"updatedAt"`
}
