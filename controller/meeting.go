package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"huiyi/dao"
	"huiyi/models"
	"net/http"
)

// 创建会议
func CreateMeeting(c *gin.Context) {
	var meet models.Meet
	if err := c.ShouldBindJSON(&meet); err != nil {
		log.Error("创建会议,绑定数据失败!", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.CreateMeeting(meet); err != nil {
		log.Error("创建会议,插入数据库失败", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	log.Info("创建会议成功")
	c.JSON(http.StatusCreated, meet)
}

// 根据id查询会议
func GetMeetings(c *gin.Context) {
	var meetid int
	if err := c.ShouldBindQuery(&meetid); err != nil {
		log.Error("会议id绑定失败!", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	meet, err := dao.GetMeeting(meetid)
	if err != nil {
		log.Error("查询数据库失败!", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	log.Info("根据id查询数据成功")
	c.JSON(http.StatusOK, meet)
}

// 更新会议
func UpdateMeeting(c *gin.Context) {
	var meet models.Meet
	if err := c.ShouldBindJSON(&meet); err != nil {
		log.Error("更新会议绑定数据失败!", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.UpdateMeeting(meet); err != nil {
		log.Error("更新会议,插入数据库失败", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	log.Info("更新会议成功")
	c.JSON(http.StatusOK, meet)
}

// 删除会议
func DeleteMeeting(c *gin.Context) {
	var meetid int
	if err := c.ShouldBindQuery(&meetid); err != nil {
		log.Error("会议id绑定失败!", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.DeleteMeeting(meetid); err != nil {
		log.Error("删除会议,数据库操作失败", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	log.Info("删除会议成功")
	c.JSON(http.StatusOK, gin.H{"deleted_id": meetid})
}

// 查询所有会议详细信息
func GetMeetingsList(c *gin.Context) {
	meetList, err := dao.GetMeetingsList()
	if err != nil {
		log.Error("查询所有会议列表失败!", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	log.Info("查询所有会议列表成功")
	c.JSON(http.StatusOK, meetList)
}
