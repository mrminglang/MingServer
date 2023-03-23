package boot

import (
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"os"
	"path/filepath"
	"server/logic"
	"server/utils/cache"
	"server/utils/confs"
	"server/utils/esdb"
	"server/utils/log"
	"server/utils/ormdb"
)

// 启动服务设置
func Boot(confNames []string, serverName string) {
	if serverName == "" {
		serverName = taf.GetServerConfig().Server
	}

	// 注册日志服务
	log.Init()
	log.Def.Infof("boot start......")

	// 注册配置信息
	confs.Init(confNames)

	// 注册MySQL服务
	_ = ormdb.Init(confs.GetConf(serverName), "db")

	// 注册DCache服务
	cache.Init(confs.GetConf(serverName))

	//初始化业务逻辑
	logic.Init(confs.GetConf(serverName))

	// 注册es服务
	esdb.Init(confs.GetConf(serverName))

	log.Def.Infof("boot success......")
}

// 关闭服务设置
func Destroy() {
	log.Destroy()
}

// 获取项目路径
func RootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
