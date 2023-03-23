package logic_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/logic"
	"server/taf-protocol/MingApp"
	"server/utils/log"
	"testing"
)

func TestMain(m *testing.M) {
	serverName := "MingServer"
	confName := boot.RootPath() + "/" + serverName
	boot.Boot([]string{confName}, confName)
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

func TestGetESDataById(t *testing.T) {
	req := new(MingApp.GetESDataByIdReq)
	rsp := new(MingApp.GetESDataByIdRsp)
	req.IndexName = "userss" //索引必须先存在
	req.Typ = "person"
	req.Id = 234
	ret, err := logic.GetESDataById(context.TODO(), req, rsp)
	if err != nil {
		log.Es.Errorf("setESData error::%s", err.Error())
		return
	}
	fmt.Printf("GetESDataById rst::%d", ret)
	fmt.Printf("GetESDataById rsp::%s", rsp.Display())
}
