package es_repository_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/repositories/es/es_repository"
	"server/taf-protocol/MingApp"
	"strconv"
	"testing"
)

var serverName = boot.RootPath() + "/MingServer"

func TestMain(m *testing.M) {
	_ = boot.Boot([]string{serverName}, serverName)
	m.Run()
}

func TestClusterHealth(t *testing.T) {
	health, err := es_repository.ClusterHealth()
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(health)
}

func TestCatIndices(t *testing.T) {
	index := "se"
	indices, err := es_repository.CatIndices(index)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(indices)
}

func TestCatAllocation(t *testing.T) {
	nodes, err := es_repository.CatAllocation()
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(nodes)
}

func TestCatMaster(t *testing.T) {
	node, err := es_repository.CatMaster()
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(node)
}

func TestIndexExists(t *testing.T) {
	indices := []string{"userss"}
	ok, err := es_repository.IndexExists(indices)
	if err != nil {
		assert.Error(t, err)
	}
	assert.True(t, ok)
}

func TestCreateIndex(t *testing.T) {
	name := "userss3"
	rsp, err := es_repository.CreateIndex(name)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(rsp)
}

func TestDeleteIndex(t *testing.T) {
	indices := []string{"userss3"}

	rsp, err := es_repository.DeleteIndex(indices)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(rsp)
}

func TestSetESData(t *testing.T) {
	teachers := MingApp.TeacherList{
		NickName:   "张三88",
		HearderPic: "http://cdn.upchinaproduct.com/project/dakaH5/images/avatar_03.png",
	}
	ret, err := es_repository.SetESData("userss", "person", strconv.Itoa(int(234)), teachers)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(ret)
}

func TestGetESDataById(t *testing.T) {
	ret, err := es_repository.GetESDataById("userss", "person", strconv.Itoa(int(234)))
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(ret)
}
