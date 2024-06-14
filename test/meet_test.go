package test

import (
	"fmt"
	"huiyi/config"
	"huiyi/dao"
	"testing"
	"time"

	"huiyi/models"
)

// 测试数据库初始化
func TestInitDB(t *testing.T) {
	// 初始化数据库
	config.InitConfigtest()
	models.InitDB()

	// 获取数据库连接
	db := models.GetDB()

	// 简单检查数据库连接是否已成功初始化
	if db == nil {
		t.Fatal("无法获取数据库连接")
	} else {
		fmt.Println("数据库连接成功")
	}
}

// 创建会议
func TestCreateMeeting(t *testing.T) {
	// 初始化测试数据库连接
	config.InitConfigtest()
	models.InitDB()

	// 创建会议
	meet := models.Meet{Name: "亲爱的同学这是一场正式的会议邀请", Time: time.Now().Add(48 * time.Hour), Location: "腾讯会议", Description: "这是一个面试", Organizer: "Dear Li"}
	err := dao.CreateMeeting(meet)
	if err != nil {
		t.Fatalf("创建会议失败: %v", err)
	}

}

// 根据id查询会议
func TestGetMeeting(t *testing.T) {
	// 初始化测试数据库连接
	config.InitConfigtest()
	models.InitDB()
	id := 7

	// 查询会议
	foundMeet, err := dao.GetMeeting(id)
	if err != nil {
		t.Fatalf("查询会议失败: %v", err)
	}
	fmt.Println(foundMeet)
}

// 更新会议
func TestUpdateMeeting(t *testing.T) {
	// 初始化测试数据库连接
	config.InitConfigtest()
	models.InitDB()

	// 创建一个会议用于测试更新
	createdMeet := models.Meet{Id: 4, Name: "亲爱的同学这是一场正式的会议邀请", Time: time.Now(), Location: "Original Location", Description: "Initial description", Organizer: "Test Organizer"}
	err := dao.UpdateMeeting(createdMeet)
	if err != nil {
		t.Fatalf("创建测试会议失败: %v", err)
	}
}

// 删除会议
func TestDeleteMeeting(t *testing.T) {
	// 初始化测试数据库连接
	config.InitConfigtest()
	models.InitDB()
	id := 4
	// 删除会议
	err := dao.DeleteMeeting(id)
	if err != nil {
		t.Fatalf("删除会议失败: %v", err)
	}
}

// 查询所有会议详细信息
func TestGetMeetingList(t *testing.T) {
	config.InitConfigtest()
	models.InitDB()
	var meets []models.Meet
	meets, err := dao.GetMeetingsList()
	if err != nil {
		t.Fatal("查询失败!!", err)
	}
	fmt.Println(meets)
}
