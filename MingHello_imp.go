package main

import (
	"context"
	"server/logic"
	"server/taf-protocol/MingApp"
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

// 设置DCache缓存
func (imp *MingHelloImp) SetStringCache(ctx context.Context, req *MingApp.SetStringCacheReq, rsp *MingApp.SetStringCacheRsp) (int32, error) {
	return logic.SetStringCache(ctx, req, rsp)
}

// 获取DCache缓存
func (imp *MingHelloImp) GetStringCache(ctx context.Context, req *MingApp.GetStringCacheReq, rsp *MingApp.GetStringCacheRsp) (int32, error) {
	return logic.GetStringCache(ctx, req, rsp)
}
