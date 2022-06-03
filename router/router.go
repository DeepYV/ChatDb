package router

import (
	"github.com/gin-gonic/gin"
	repositories "test.com/test/repository"
	services "test.com/test/service"
)

func Router(router *gin.Engine) {

	userDBClient := repositories.NewRepository()
	usercontroller := services.NewService(userDBClient)

	router.POST(CREATEUSER, usercontroller.CreateUser)
	router.POST(CREATEGROUP, usercontroller.CreateChatRoom)
	router.POST(CREATEMESSAGE, usercontroller.CreateMessage)
	router.POST(JOINGROUP, usercontroller.JoinedGroup)
	router.GET(GETCHATROOM, usercontroller.GetAllChatRoom)
	router.DELETE(CHATROOM, usercontroller.DeleteChatRoom)
	return
}
