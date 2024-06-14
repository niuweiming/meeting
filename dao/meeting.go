package dao

import (
	log "github.com/sirupsen/logrus"
	"huiyi/models"
)

// 创建会议
func CreateMeeting(meet models.Meet) error {
	db := models.DB
	err := db.Create(&meet).Error
	if err != nil {
		log.Error("插入会议失败: ", err)
		return err
	}
	log.Info("插入会议成功: ", meet)
	return nil
}

// 根据id查询会议
func GetMeeting(meetid int) (models.Meet, error) {
	db := models.DB
	var meet models.Meet
	err := db.Where("id = ?", meetid).First(&meet).Error
	if err != nil {
		log.Error("查询会议失败: ", err)
		return models.Meet{}, err
	}
	log.Info("查询会议成功: ", meet)
	return meet, nil
}

// 删除会议
func DeleteMeeting(meetId int) error {
	db := models.DB
	err := db.Where("id = ?", meetId).Delete(&models.Meet{}).Error
	if err != nil {
		log.Error("删除会议失败: ", err)
		return err
	}
	log.Info("删除会议成功, ID: ", meetId)
	return nil
}

// 更新会议
func UpdateMeeting(meet models.Meet) error {
	db := models.DB
	err := db.Model(&meet).Updates(&meet).Error
	if err != nil {
		log.Error("更新会议失败: ", err)
		return err
	}
	log.Info("更新会议成功: ", meet)
	return nil
}

// 查询所有会议详细信息
func GetMeetingsList() ([]models.Meet, error) {
	var meets []models.Meet
	db := models.DB
	err := db.Find(&meets).Error
	if err != nil {
		log.Error("查询所有会议数据出错: ", err)
		return nil, err
	}
	log.Info("查询所有会议数据成功: ", meets)
	return meets, nil
}
