package main

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"server/taf-protocol/MingApp"
	"server/utils/log"
)

func getTeacherList(app *MingApp.MingHello) {
	req := new(MingApp.GetTeacherListReq)
	rsp := new(MingApp.GetTeacherListRsp)
	req.BeginIndex = 0
	req.Count = 10
	//req.NickName = "张三"
	ret, err := app.GetTeacherList(req, rsp)
	if err != nil {
		log.Def.Errorf("getTeacherList error::%s", err.Error())
		return
	}
	fmt.Printf("getTeacherList rst::%d", ret)
	fmt.Printf("getTeacherList rsp::%s", rsp.Display())
}

func setESData(app *MingApp.MingHello) {
	req := new(MingApp.SetESDataReq)
	rsp := new(MingApp.SetESDataRsp)
	req.IndexName = "userss"
	req.Typ = "person"
	req.Id = 8
	req.Teachers = MingApp.TeacherList{
		NickName:   "张三8",
		HearderPic: "http://cdn.upchinaproduct.com/project/dakaH5/images/avatar_03.png",
	}
	ret, err := app.SetESData(req, rsp)
	if err != nil {
		log.Es.Errorf("setESData error::%s", err.Error())
		return
	}
	fmt.Printf("getTeacherList rst::%d", ret)
	fmt.Printf("getTeacherList rsp::%s", rsp.Display())
}

func getESDataById(app *MingApp.MingHello) {
	req := new(MingApp.GetESDataByIdReq)
	rsp := new(MingApp.GetESDataByIdRsp)
	req.IndexName = "userss" //索引必须先存在
	req.Typ = "person"
	req.Id = 8
	ret, err := app.GetESDataById(req, rsp)
	if err != nil {
		log.Es.Errorf("setESData error::%s", err.Error())
		return
	}
	fmt.Printf("getTeacherList rst::%d", ret)
	fmt.Printf("getTeacherList rsp::%s", rsp.Display())
}

func main() {
	comm := taf.NewCommunicator()
	obj := fmt.Sprintf("MingApp.MingServer.MingHelloObj@tcp -h 127.0.0.1 -p 20230 -t 60000")
	app := new(MingApp.MingHello)
	comm.StringToProxy(obj, app)

	//getTeacherList(app)
	setESData(app)
	getESDataById(app)
}
