package boot

import (
	"server/logic"
	"server/utils/cache"
	"server/utils/conf"
	"server/utils/esdb"
	"server/utils/log"
	"server/utils/ormdb"
)

// 启动服务设置
func Boot() {
	// 注册日志服务
	log.Init()
	log.Def.Infof("boot start......")

	// 注册配置信息
	conf.Init()

	// 注册MySQL服务
	_ = ormdb.Init(conf.GetConf(), "db")

	// 注册DCache服务
	cache.Init(conf.GetConf())

	// 注册es服务
	esdb.Init(conf.GetConf())

	//初始化业务逻辑
	logic.Init(conf.GetConf())

	log.Def.Infof("boot success......")
}

// 关闭服务设置
func Destroy() {
	log.Destroy()
}
