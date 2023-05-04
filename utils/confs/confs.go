package confs

import (
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"gitlab.upchinaproduct.com/upgo/utils/log"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/taftool"
)

var ConfName2Ptr = make(map[string]*conf.Conf)

// 初始化配置
func Init(confNames []string) error {
	log.Conf.Infof("init config start......")
	cfg := taf.GetServerConfig()
	if len(confNames) <= 0 {
		confNames = append(confNames, cfg.Server)
	}

	log.Conf.Infof("init config confNames::", confNames)

	for _, name := range confNames {
		log.Conf.Infof("init config for start name::", name)
		confPtr, err := taftool.GetTafConfigByName(name + ".conf")
		if err != nil {
			log.Conf.Errorf("init config for name:: error::", name, err.Error())
			return err
		}
		if confPtr == nil {
			log.Conf.Errorf("init config for is null name::", name)
			return err
		}
		ConfName2Ptr[name] = confPtr
		log.Conf.Infof("init config for success name::", name)
	}

	log.Conf.Infof("init config success......", ConfName2Ptr)
	return nil
}

// 获取配置
func GetConf(confName string) *conf.Conf {
	return ConfName2Ptr[confName]
}
