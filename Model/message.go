package Model

const (
	TABLE_MESSAGE                     = "message"
	TABLE_MESSAGE_COLUMN_USER_ID      = "user_id"
	TABLE_MESSAGE_COLUMN_CREATEDAT    = "createdat"
	TABLE_MESSAGE_COLUMN_MESSSAGE     = "message"
	TABLE_MESSAGE_COLUMN_CHATROOMNAME = "chatroom"
)

type Message struct {
	User_id      *string `json:"user_id"`
	Createdat   string `json:"createdat"`
	Chatroomname *string `json:"chatroomname"`
	Message      *string `json:"message"`
}
