package esrpc_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/logic/esrpc"
	"server/taf-protocol/FCS"
	"testing"
)

func TestMain(m *testing.M) {
	// 绝对路径
	confName := "/Users/ming/work/up/MingApp/MingServer/MingServer"
	boot.Boot([]string{confName})
	m.Run()
}

func TestGetESClusterList(t *testing.T) {
	esReq := FCS.GetESClusterListReq{
		IsCluster: true,
	}
	esRsp := FCS.GetESClusterListRsp{}
	var err error
	esRsp, err = esrpc.GetESClusterList(esReq)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(esRsp)
}
