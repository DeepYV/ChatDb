package Model

const (
	TABLE_USER                 = "user"
	TABLE_USER_COLUMN_USER_ID  = "user_id"
	TABLE_USER_COLUMN_USERNAME = "username"
	TABLE_USER_COLUMN_PASSWORD = "password"
	TABLE_USER_COLUMN_ONLINE   = "online"
	TABLE_USER_COLUMN_CREATEAT = "createat"
)

type User struct {
	User_id  string `json:"user_id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Online   *string `json:"online"`
	Createat string `json:"createat"`
}
