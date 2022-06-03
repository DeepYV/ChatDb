package Model

const (
	TABLE_CHATROOM                  = "chatroom"
	TABLE_CHATROOM_COLUMN_CHATROOM  = "chatroomname"
	TABLE_CHATROOM_COLUMN_CREATEDAT = "createdat"
	TABLE_CHATROOM_COLUMN_MEMBER    = "member"
)

type ChatRoom struct {
	Chatroomname *string `json:"chatroom"`
	Createdat   string  `json:"createdat"`
	Member       int     `json:"member"`
}
