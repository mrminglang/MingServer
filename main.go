package main

import (
	"fmt"
	"os"
	"server/boot"
	"server/logic/http_proxy"
	"server/taf-protocol/MingApp"

	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/confs"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/log"
)

func main() {
	// Get server config
	cfg := taf.GetServerConfig()

	// 启动boot
	err := boot.Boot([]string{}, cfg.Server)
	if err != nil {
		fmt.Printf("MingHelloImp init boot, err:(%s)\n", err.Error())
		os.Exit(-1)
	}

	// New servant imp
	{
		imp := new(MingHelloImp)
		err = imp.Init()
		if err != nil {
			fmt.Printf("MingHelloImp init fail, err:(%s)\n", err.Error())
			os.Exit(-1)
		}
		app := new(MingApp.MingHello)                                          // New servant
		app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".MingHelloObj") // Register Servant
	}

	// http
	{
		if confs.GetConf(cfg.Server).GetBoolWithDef("/app/<isHttpOn>", false) {
			// http接口监听 没走taf协议，请求不被taf接管
			log.Def.Infof("MingHelloImp init http start......")
			httpProxy := http_proxy.GetHttpProxy(cfg.App + "." + cfg.Server + ".MingHttpObj")
			httpProxy.Run()
		}
	}

	// Run application
	taf.Run()
}
