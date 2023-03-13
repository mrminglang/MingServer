package conf

import (
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"server/utils/log"
	"server/utils/taftool"
)

var Conf *conf.Conf

// 初始化配置
func Init(confName string) {
	log.Conf.Infof("init config start......")
	log.Conf.Infof("init config confName::", confName)
	confs, err := taftool.GetTafConfigByName(confName)
	if err != nil {
		log.Conf.Errorf("init config error::", err.Error())
		return
	}
	if confs == nil {
		log.Conf.Errorf("init config is null")
		return
	}
	Conf = confs
	log.Conf.Infof("init config success......")
}

// 获取配置
func GetConf() *conf.Conf {
	return Conf
}
