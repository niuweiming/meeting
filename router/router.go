package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"huiyi/controller"
)

func RouterStart() {
	router := gin.Default()
	//会议
	{
		router.POST("/meetings", controller.CreateMeeting)
		router.GET("/meetings", controller.GetMeetings)
		router.PUT("/meetings/:id", controller.UpdateMeeting)
		router.DELETE("/meetings/:id", controller.DeleteMeeting)
		router.GET("/meetings/all", controller.GetMeetingsList)
	}
	//邀请
	{
		router.POST("/invites", controller.CreateInvite)
		router.GET("/invites", controller.GetInvites)
		router.PUT("/invites/:id", controller.UpdateInvite)
	}

	log.Info(router.Run(":8080"))
}
