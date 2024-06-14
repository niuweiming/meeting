package app

import (
	log "github.com/sirupsen/logrus"
	"huiyi/config"
	"huiyi/models"
	"huiyi/router"
)

func Start() {
	// 初始化配置
	config.InitConfig()
	models.InitDB()
	router.RouterStart()
	log.Error("启动成功")
}
