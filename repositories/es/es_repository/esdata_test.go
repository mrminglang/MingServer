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

func TestMain(m *testing.M) {
	// 绝对路径
	confName := "/Users/ming/work/up/MingApp/MingServer/MingServer.conf"
	boot.Boot(confName)
	m.Run()
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
