package trpc

import (
	"fmt"
	"gitlab.upchinaproduct.com/upgo/utils/log"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/taftool"
	"server/taf-protocol/FCS"
)

var ESPrx FCS.ESDriver

func ESInit(obj string) error {
	//初始化esRpc链接
	log.Es.Infof(fmt.Sprintf("init esrpc server obj::%s start....", obj))
	taftool.InitTafPrx(&ESPrx, obj)
	log.Es.Infof("init esrpc server success....")

	return nil
}
