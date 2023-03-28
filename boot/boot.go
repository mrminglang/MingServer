package boot

import (
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"os"
	"path/filepath"
	"server/logic"
	"server/utils/confs"
	"server/utils/esdb"
	"server/utils/log"
	"server/utils/ormdb"
)

// 启动服务设置
func Boot(confNames []string, serverName string) error {
	if serverName == "" {
		serverName = taf.GetServerConfig().Server
	}

	// 注册日志服务
	log.Init()
	log.Def.Infof("boot start......")

	// 注册配置信息
	err := confs.Init(confNames)
	if err != nil {
		log.Def.Errorf("boot confs error::", err.Error())
		return err
	}

	// 注册MySQL服务
	err = ormdb.Init(confs.GetConf(serverName), "db")
	if err != nil {
		log.Def.Errorf("boot ormdb error::", err.Error())
		return err
	}

	//初始化业务逻辑
	err = logic.Init(confs.GetConf(serverName))
	if err != nil {
		log.Def.Errorf("boot logic error::", err.Error())
		return err
	}

	// 注册es服务
	err = esdb.Init(confs.GetConf(serverName))
	if err != nil {
		log.Def.Errorf("boot esdb error::", err.Error())
		return err
	}

	log.Def.Infof("boot success......")
	return nil
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
