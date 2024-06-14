package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"huiyi/dao"
	"huiyi/models"
	"huiyi/utils"
	"net/http"
)

// 创建邀请
func CreateInvite(c *gin.Context) {
	var invite models.Invite
	if err := c.ShouldBindJSON(&invite); err != nil {
		log.Error("绑定数据失败: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.CreateInvite(invite); err != nil {
		log.Error("创建邀请失败: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//发送邮件
	if err := Mail(invite); err != nil {
		log.Error("邮件发送失败: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("创建邀请成功: ", invite)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": invite})
}

// 查询邀请
func GetInvites(c *gin.Context) {
	var inviteId int
	if err := c.ShouldBindQuery(&inviteId); err != nil {
		log.Error("绑定数据失败: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invite, err := dao.GetInvite(inviteId)
	if err != nil {
		log.Error("查询邀请失败: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("查询邀请成功: inviteId=", inviteId)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": invite})
}

// 更新邀请
func UpdateInvite(c *gin.Context) {
	var invite models.Invite
	if err := c.ShouldBindJSON(&invite); err != nil {
		log.Error("绑定数据失败: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.UpdataInvite(invite); err != nil {
		log.Error("更新邀请失败: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("更新邀请成功: ", invite)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": invite})
}

// 发送邮件
func Mail(invite models.Invite) error {
	contect, err := dao.GetMeeting(invite.MeetingId)
	if err != nil {
		log.Error("获取会议信息失败: ", err)
		return err
	}

	body := fmt.Sprintf("会议内容为:\n名称：%s\n时间：%s\n地点：%s\n描述：%s\n组织者：%s",
		contect.Name,
		contect.Time.Format("2006-01-02 15:04"),
		contect.Location,
		contect.Description,
		contect.Organizer,
	)

	mailConf := models.MailboxConf{
		Title:         "您有一场会议待查收!",
		Body:          body,
		RecipientList: []string{invite.Email}, // 收件人邮箱地址
		Sender:        "hkwl_0514@qq.com",     // 发件人的QQ邮箱地址
		SPassword:     "jpnrymxonnkyddgi",     // 你的QQ邮箱SMTP授权码
		SMTPAddr:      "smtp.qq.com",          // QQ邮箱SMTP服务器地址
		SMTPPort:      465,                    // QQ邮箱SMTP端口，465为SSL加密端口，587为TLS加密端口
	}

	// 使用gomail发送邮件
	err = utils.SendMail(mailConf)
	if err != nil {
		log.Error("发送邮件失败: ", err)
		return err
	}
	log.Info("邮件发送成功")
	return nil
}
