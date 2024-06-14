package dao

import (
	log "github.com/sirupsen/logrus"
	"huiyi/models"
)

// 创建邀请
func CreateInvite(invite models.Invite) error {
	db := models.DB
	// 尝试将邀请记录插入到数据库中
	err := db.Create(&invite).Error
	if err != nil {
		log.Error("插入邀请失败: ", err)
		return err
	}
	log.Info("插入邀请成功: ", invite)
	return nil
}

// 根据邀请ID查询邀请记录
func GetInvite(inviteid int) (models.Invite, error) {
	db := models.DB
	var invite models.Invite
	// 尝试根据ID从数据库中查询邀请记录
	err := db.First(&invite, inviteid).Error
	if err != nil {
		log.Error("查询邀请失败: ", err)
		return models.Invite{}, err
	}
	log.Info("查询邀请成功: ", invite)
	return invite, nil
}

// 更新邀请记录 主要是状态更新
func UpdataInvite(invite models.Invite) error {
	db := models.DB
	// 尝试更新邀请记录
	err := db.Model(&invite).Updates(invite).Error
	if err != nil {
		log.Error("邀请更新失败: ", err)
		return err
	}
	log.Info("邀请更新成功: ", invite)
	return nil
}
