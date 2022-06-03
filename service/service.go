package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	model "test.com/test/Model"
	repositories "test.com/test/repository"
	"test.com/test/request"
)

type IService interface {
	CreateUser(ctx *gin.Context) error
	CreateChatRoom(ctx *gin.Context) error
	CreateMessage(ctx *gin.Context) error
	Joingroup(ctx gin.Context) error
	GetAllGroup(ctx *gin.Context) error
	GetAllChatRoom(ctx *gin.Context) error
	DeleteChatRoom(ctx *gin.Context, Chatroom *string) error
}

type Service struct {
	service repositories.IRepository
}

func NewService(servicedata repositories.IRepository) *Service {
	return &Service{service: servicedata}

}
func (u *Service) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		log.Fatalln("error while decoding request body", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Wrong body "})
		return
	}
	uid, _ := uuid.NewUUID()
	user.User_id = uid.String()

	user.Createat = time.Now().Format(time.RFC3339)
	err := u.service.CreateUser(ctx, user)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"MESSAGE": "user created"})
	return
}

func (u *Service) CreateChatRoom(ctx *gin.Context) {

	var Chatroom model.ChatRoom
	if err := json.NewDecoder(ctx.Request.Body).Decode(&Chatroom); err != nil {
		log.Fatalln("error while decoding request body", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Wrong body "})
		return
	}
	err := u.service.CreateChatRoom(ctx, Chatroom)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"MESSAGE": "Chatroom created"})
	return

}
func (u *Service) CreateMessage(ctx *gin.Context) {

	var message model.Message
	if err := json.NewDecoder(ctx.Request.Body).Decode(&message); err != nil {
		log.Fatalln("error while decoding request body", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Wrong body "})
		return
	}
	err := u.service.CreateMessageRoom(ctx, message)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"MESSAGE": "Message created"})
	return

}
func (u *Service) JoinedGroup(ctx *gin.Context) {

	var Joined model.Joined
	if err := json.NewDecoder(ctx.Request.Body).Decode(&Joined); err != nil {
		log.Fatalln("error while decoding request body", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Wrong body "})
		return
	}
	err := u.service.JoinRoom(ctx, Joined)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"MESSAGE": "User Added tp the group"})
	return

}
func (u *Service) GetAllChatRoom(ctx *gin.Context) {

	Chatroom, err := u.service.GetAllChatRoom(ctx)
	if err != nil {

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return
	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"All ChatRoom": Chatroom})
	return

}

func (u *Service) DeleteChatRoom(ctx *gin.Context) {

	var Chatroom request.DeleteChatRoom
	if err := json.NewDecoder(ctx.Request.Body).Decode(&Chatroom); err != nil {
		log.Fatalln("error while decoding request body", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Wrong body "})
		return
	}
	err := u.service.DeleteChatRoom(ctx, Chatroom)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Retry again"})
		return

	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"ChatRoom deleted": Chatroom.Chatroomname})
	return

}
