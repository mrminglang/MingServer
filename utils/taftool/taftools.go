package taftool

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
)

var tafConf = taf.GetServerConfig()

//获取taf配置
func GetTafConfigByName(confName string) (*conf.Conf, error) {
	if !IsLocalEnv() {
		//在taf环境工作
		//从配置中心定位配置
		remoteConf := taf.NewRConf(tafConf.App, tafConf.Server, tafConf.BasePath)
		//获取配置文件到本地
		remoteConfContent, _ := remoteConf.GetConfig(confName)
		fmt.Printf("{getRemoteConf confName^content}|%s|%s\n", confName, remoteConfContent)
		confName = tafConf.BasePath + "/" + confName
	}
	conf, err := conf.NewConf(confName)
	contentParse := conf.ToString()
	fmt.Println("{parseConfRes}", contentParse)
	if err != nil {
		fmt.Printf("{newConfError confName^err}|%s|%s\n", confName, err.Error())
	}

	return conf, err
}

//返回 是否在本地环境
func IsLocalEnv() bool {
	return tafConf.BasePath == ""
}

//链接taf接口
func InitTafPrx(prx taf.ProxyPrx, obj string) {
	comm := taf.NewCommunicator()
	comm.StringToProxy(obj, prx)
}
