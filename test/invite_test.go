package test

import (
	"fmt"
	"huiyi/config"
	"huiyi/controller"
	"huiyi/dao"
	"huiyi/models"
	"testing"
	"time"
)

// 创建邀请
func TestCreateInvite(t *testing.T) {
	// 初始化测试数据库连接
	config.InitConfigtest()
	models.InitDB()

	// 准备待插入的邀请信息
	invite := models.Invite{
		MeetingId: 7,
		Email:     "1214602074@qq.com",
		SentTime:  time.Now(),
		Status:    "sent",
	}

	// 插入邀请
	err := dao.CreateInvite(invite)
	if err != nil {
		t.Fatalf("插入邀请失败: %v", err)
	}

}

// 查询邀请
func TestGetInvite(t *testing.T) {
	// 初始化测试数据库连接
	config.InitConfigtest()
	models.InitDB()
	id := 6
	// 查询邀请
	fetchedInvite, err := dao.GetInvite(id)
	if err != nil {
		t.Fatalf("查询邀请失败: %v", err)
	}
	fmt.Println(fetchedInvite)

}

// 更新邀请
func TestUpdateInvite(t *testing.T) {
	// 初始化测试数据库连接
	config.InitConfigtest()
	models.InitDB()
	id := 4
	// 更新邀请信息
	updatedInfo := models.Invite{
		Id:     id,
		Status: "accepted",
	}
	err := dao.UpdataInvite(updatedInfo)
	if err != nil {
		t.Fatalf("更新邀请失败: %v", err)
	}
}

// 邮件发送
func TestSendemailInvite(t *testing.T) {
	config.InitConfigtest()
	models.InitDB()
	invite := models.Invite{
		MeetingId: 7,
		Email:     "1214602074@qq.com",
		SentTime:  time.Now(),
		Status:    "sent",
	}

	err := controller.Mail(invite)
	if err != nil {
		t.Fatalf("发送失败: %v", err)
	}
	fmt.Println("发送成功!!")

}
