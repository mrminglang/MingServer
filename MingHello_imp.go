package main

import (
	"MingServer/logic"
	"MingServer/taf-protocol/MingApp"
	"context"
)

// MingHelloImp servant implementation
type MingHelloImp struct {
}

// Init servant init
func (imp *MingHelloImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (imp *MingHelloImp) Destroy() {
	//destroy servant here:
	//...
}

// 获取老师列表
func (imp *MingHelloImp) GetTeacherList(ctx context.Context, req *MingApp.GetTeacherListReq, rsp *MingApp.GetTeacherListRsp) (int32, error) {
	return logic.GetTeacherList(ctx, req, rsp)
}
