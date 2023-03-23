package teacher_repository_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/repositories/mysql/teacher_repository"
	"testing"
)

func TestMain(m *testing.M) {
	serverName := "MingServer"
	confName := boot.RootPath() + "/" + serverName
	boot.Boot([]string{confName}, confName)
	m.Run()
}

func TestQueryTeachers(t *testing.T) {
	whereMaps := map[string]string{
		"nickname": "张三",
		"order":    "createtime ASC",
	}

	total, teachers, err := teacher_repository.QueryTeachers(0, 10, whereMaps)
	assert.Nil(t, err)
	dumps.Dump(total)
	dumps.Dump(teachers)

	boot.Destroy()
}
