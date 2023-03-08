package conf

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"server/utils/taftool"
)

var Conf *conf.Conf

// 初始化配置
func InitConf() {
	fmt.Printf("init config start ......\n")
	cfg := taf.GetServerConfig()
	fmt.Printf("init config Server::%s\n", cfg.Server)
	confName := cfg.Server + ".conf"
	fmt.Printf("init config confName::::%s\n", confName)
	confs, err := taftool.GetTafConfigByName(confName)
	if err != nil {
		_ = fmt.Errorf("init config error::%s\n", err.Error())
		return
	}
	if confs == nil {
		_ = fmt.Errorf("init config is null")
		return
	}
	Conf = confs
	fmt.Printf("init config success......\n")
}

// 获取配置
func GetConf() *conf.Conf {
	return Conf
}
