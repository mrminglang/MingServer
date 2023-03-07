package boot

import (
	"MingServer/utils/cache"
	"MingServer/utils/conf"
	"MingServer/utils/log"
	"MingServer/utils/ormdb"
)

// 启动服务设置
func Boot() {
	// 注册日志服务
	log.InitLog()
	log.Def.Infof("boot start......")

	// 注册配置信息
	conf.InitConf()

	// 注册MySQL服务
	_ = ormdb.InitDb(conf.GetConf(), "db")

	// 注册DCache服务
	cache.Init(conf.GetConf())

	log.Def.Infof("boot success......")
}

// 关闭服务设置
func Destroy() {
	log.DestroyLog()
}
