package main

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"os"
	"server/boot"
	"server/taf-protocol/MingApp"
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

	// 支持HTTP
	//{
	//	mux := &taf.TafHttpMux{}
	//	mux.HandleFunc("/ming", func(w http.ResponseWriter, r *http.Request) {
	//		whereMaps := map[string]string{
	//			//"nickname": "张三",
	//			"order": "createtime ASC",
	//		}
	//		_, teachers, _ := teacher_repository.NewTeacher().QueryTeachers(0, 100, whereMaps)
	//		resp, _ := json.Marshal(teachers)
	//
	//		_, _ = w.Write(resp)
	//	})
	//	taf.AddHttpServant(mux, cfg.App+"."+cfg.Server+".MingHelloObj") //Register http server
	//}

	// Run application
	taf.Run()
}
