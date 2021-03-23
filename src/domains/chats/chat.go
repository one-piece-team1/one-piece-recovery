package chats

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/chatparticipates"
)

type ChatSendStatus = string

// Conflict with type type move to here
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

type ChatDomains struct {
	ID              string                                  `json:"id"`
	Name            string                                  `json:"name"`
	SendStatus      ChatSendStatus                          `json:"sendStatus"`
	ChatParticipate chatparticipates.ChatParitcipateDomains `json:"chatParticipate"`
	CreatedAt       string                                  `json:"createdAt"`
	UpdatedAt       string                                  `json:"updatedAt"`
}
