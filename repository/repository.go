package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	database "test.com/test/Database"
	model "test.com/test/Model"
	"test.com/test/request"
)

type IRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	CreateChatRoom(ctx context.Context, chatroom model.ChatRoom) error
	CreateMessageRoom(ctx context.Context, message model.Message) error
	JoinRoom(ctx context.Context, Joined model.Joined) error
	GetAllChatRoom(ctx context.Context) ([]string, error)
	DeleteChatRoom(ctx context.Context, Chatroom request.DeleteChatRoom) error
}

type db struct {
	db *gorm.DB
}

func NewRepository() *db {
	return &db{
		db: database.GetConnection()}
}
func (conn *db) CreateUser(ctx context.Context, user model.User) error {

	err := conn.db.Table(model.TABLE_USER).Create(&user)
	if err.Error != nil {
		return err.Error

	}

	return nil
}

// user login --> create room --> extract uiser_id from context and marked him as ADMIN
func (conn *db) CreateChatRoom(ctx context.Context, chatroom model.ChatRoom) error {

	chatroom.Createdat = time.Now().Format(time.RFC3339)
	chatroom.Member = 1

	err := conn.db.Table(model.TABLE_CHATROOM).Create(&chatroom)
	if err.Error != nil {
		return err.Error
	}
	query := fmt.Sprintf("INSERT INTO %s (%s,%s,%s, %s,%s) VALUES ('%s','%s','%s','%s','%s')",
		model.TABLE_JOINED,
		model.TABLE_JOINED_COLUMN_JOIN_ID,
		model.TABLE_JOINED_COLUMN_USER_ID,
		model.TABLE_JOINED_COLUMN_CHATROOMNAME,
		model.TABLE_JOINED_COLUMN_CREATEDAT,
		model.TABLE_JOINED_COLUMN_ROLE,
		uuid.New().String(),
		"1",
		*chatroom.Chatroomname,
		chatroom.Createdat,
		model.ADMIN)

	if err = conn.db.Exec(query); err.Error != nil {

		return err.Error
	}
	return nil
}

// create message and refer with userid
func (conn *db) CreateMessageRoom(ctx context.Context, message model.Message) error {

	message.Createdat = time.Now().Format(time.RFC3339)
	err := conn.db.Table(model.TABLE_MESSAGE).Create(&message)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
func (conn *db) JoinRoom(ctx context.Context, Joined model.Joined) error {
	Joined.Createdat = time.Now().Format(time.RFC3339)
	Joined.Join_id = uuid.New().String()
	query := fmt.Sprintf("INSERT INTO %s (%s,%s,%s, %s,%s) VALUES ('%s','%s','%s','%s', '%s')",
		model.TABLE_JOINED,
		model.TABLE_JOINED_COLUMN_JOIN_ID,
		model.TABLE_JOINED_COLUMN_USER_ID,
		model.TABLE_JOINED_COLUMN_CHATROOMNAME,
		model.TABLE_JOINED_COLUMN_CREATEDAT,
		model.TABLE_JOINED_COLUMN_ROLE,
		uuid.New(),
		*Joined.User_id,
		*Joined.Chatroomname,
		Joined.Createdat,
		model.MEMBER)

	if err := conn.db.Exec(query); err.Error != nil {

		return err.Error
	}
	return nil

}

func (conn *db) GetAllChatRoom(ctx context.Context) ([]string, error) {

	rows, err := conn.db.Table(model.TABLE_CHATROOM).Select(model.TABLE_CHATROOM_COLUMN_CHATROOM).Rows()

	var chatroom []string
	if err != nil {
		return []string{}, err
	}
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return []string{}, err
		}
		chatroom = append(chatroom, s)
	}
	return chatroom, nil
}
func (conn *db) DeleteChatRoom(ctx context.Context, Chatroom request.DeleteChatRoom) error {



	query := fmt.Sprintf( "DELETE FROM  %s WHERE CHATROOMNAME = '%s'", model.TABLE_CHATROOM, *Chatroom.Chatroomname)


	if err := conn.db.Exec(query); err.Error != nil {

		return err.Error
	}
	return nil
}
