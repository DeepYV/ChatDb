package Model

const (
	ADMIN                            = "admin"
	MEMBER                           = "member"
	TABLE_JOINED                     = "joined"
	TABLE_JOINED_COLUMN_USER_ID      = "user_id"
	TABLE_JOINED_COLUMN_JOIN_ID      = "join_id"
	TABLE_JOINED_COLUMN_CREATEDAT    = "createdat"
	TABLE_JOINED_COLUMN_DELETEDAT    = "deletedat"
	TABLE_JOINED_COLUMN_ROLE         = "role"
	TABLE_JOINED_COLUMN_CHATROOMNAME = "chatroomname"
)

type Joined struct {
	User_id      *string `json:"user_id"`
	Join_id      string  `json:"join_id "`
	Createdat   string  `json:"createdat"`
	Deletedat    *string `json:"deletedat"`
	Chatroomname *string `json:"chatroomname"`
	Role         *string `json:"role"`
}
