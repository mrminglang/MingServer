package main

import (
	"MingServer/taf-protocol/MingApp"
	"MingServer/utils/log"
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
)

func getTeacherList(app *MingApp.MingHello) {
	req := new(MingApp.GetTeacherListReq)
	rsp := new(MingApp.GetTeacherListRsp)
	req.BeginIndex = 0
	req.Count = 10
	//req.NickName = "张三"
	rst, err := app.GetTeacherList(req, rsp)
	if err != nil {
		log.Def.Errorf("getTeacherList error::%s", err.Error())
		return
	}
	fmt.Printf("getTeacherList rst::%d", rst)
	fmt.Printf("getTeacherList rsp::%s", rsp.Display())
}

func main() {
	comm := taf.NewCommunicator()
	obj := fmt.Sprintf("MingApp.MingServer.MingHelloObj@tcp -h 127.0.0.1 -p 20230 -t 60000")
	app := new(MingApp.MingHello)
	comm.StringToProxy(obj, app)

	getTeacherList(app)
}
