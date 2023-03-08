package esrpc

import (
	"server/taf-protocol/FCS"
	"server/utils/log"
	"server/utils/taftool"
)

var esPrx FCS.ESDriver

func Init(esObj string) {
	//初始化esrpc链接
	log.Es.Infof("init broker server start....")
	taftool.InitTafPrx(&esPrx, esObj)
	log.Es.Infof("init broker server success....")
}
