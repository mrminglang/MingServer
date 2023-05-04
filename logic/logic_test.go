package logic_test

import (
	"context"
	"fmt"
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/log"
	"server/boot"
	"server/logic"
	"server/taf-protocol/MingApp"
	"testing"
)

var serverName = boot.RootPath() + "/MingServer"

func TestMain(m *testing.M) {
	_ = boot.Boot([]string{serverName}, serverName)
	m.Run()
}

func TestGetTeacherList(t *testing.T) {
	req := new(MingApp.GetTeacherListReq)
	rsp := new(MingApp.GetTeacherListRsp)
	req.BeginIndex = 0
	req.Count = 10
	req.NickName = "张三"
	ret, err := logic.GetTeacherList(context.TODO(), req, rsp)
	if err != nil {
		assert.Error(t, err)
		return
	}
	fmt.Printf("getTeacherList rst::%d", ret)
	fmt.Printf("getTeacherList rsp::%s", rsp.Display())
}

func TestSetESData(t *testing.T) {
	req := new(MingApp.SetESDataReq)
	rsp := new(MingApp.SetESDataRsp)
	req.IndexName = "users"
	req.Typ = "person"
	req.Id = 2
	req.Teachers = MingApp.TeacherList{
		NickName:   "张三三",
		HearderPic: "http://cdn.upchinaproduct.com/project/dakaH5/images/avatar_03.png",
	}

	ret, err := logic.SetESData(context.TODO(), req, rsp)
	if err != nil {
		assert.Error(t, err)
		return
	}

	dumps.Dump(ret)

	boot.Destroy()
}

func TestGetESDataById(t *testing.T) {
	req := new(MingApp.GetESDataByIdReq)
	rsp := new(MingApp.GetESDataByIdRsp)
	req.IndexName = "users" //索引必须先存在
	req.Typ = "person"
	req.Id = 2
	ret, err := logic.GetESDataById(context.TODO(), req, rsp)
	if err != nil {
		log.Es.Errorf("setESData error::%s", err.Error())
		return
	}
	fmt.Printf("GetESDataById rst::%d", ret)
	fmt.Printf("GetESDataById rsp::%s", rsp.Display())

	boot.Destroy()
}
