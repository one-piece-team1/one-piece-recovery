package chatparticipates

import (
	"github.com/one-piece-team1/one-piece-recovery/src/domains/chatrooms"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/chats"
	"github.com/one-piece-team1/one-piece-recovery/src/domains/users"
)

type ChatParitcipateDomains struct {
	Id        string                    `json:"id"`
	ChatRoom  chatrooms.ChatRoomDomains `json:"chatroom"`
	Chats     []chats.ChatDomains       `json:"chats"`
	Users     []users.UserDomains       `json:"users"`
	CreatedAt string                    `json:"createdAt"`
	UpdatedAt string                    `json:"updatedAt"`
}
