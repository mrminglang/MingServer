package boot

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"gitlab.upchinaproduct.com/upgo/utils/esdb"
	"gitlab.upchinaproduct.com/upgo/utils/log"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/ormdb"
	"os"
	"path/filepath"
	"server/logic"
	"server/logic/esrpc"
	"server/taf-protocol/FCS"
	"server/utils/confs"
	"server/utils/trpc"
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
	err = ormdb.InitDb(confs.GetConf(serverName), "db")
	if err != nil {
		log.Def.Errorf("boot ormdb error::", err.Error())
		return err
	}

	// ESDriverServer rpc
	err = trpc.ESInit(confs.GetConf(serverName).GetString("/obj/<esObj>"))
	if err != nil {
		log.Def.Infof("{boot esinit error::}", err.Error())
		return err
	}

	// 获取ES集群配置
	esRsp, err := esrpc.GetESClusterList(FCS.GetESClusterListReq{})
	if err != nil {
		log.Es.Errorf("{esdb init GetESClusterList error|%s}", err.Error())
		return err
	}

	hosts := make([]string, 0)
	for _, row := range esRsp.EsClusters {
		http := fmt.Sprintf("http://%s:%s", row.Host, row.Port)
		hosts = append(hosts, http)
	}

	// 注册es服务
	user := confs.GetConf(serverName).GetString("/esConf/<username>")
	password := confs.GetConf(serverName).GetString("/esConf/<password>")
	err = esdb.Init(hosts, user, password)
	if err != nil {
		log.Def.Errorf("boot esdb error::", err.Error())
		return err
	}

	// 注册DCache服务
	err = trpc.DCacheInit(confs.GetConf(serverName), "dcache")
	if err != nil {
		log.Def.Errorf("boot DCache error::", err.Error())
		return err
	}

	//初始化业务逻辑
	err = logic.Init(confs.GetConf(serverName))
	if err != nil {
		log.Def.Errorf("boot logic error::", err.Error())
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
