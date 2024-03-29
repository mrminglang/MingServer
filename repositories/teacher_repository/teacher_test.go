package teacher_repository_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/repositories/teacher_repository"
	"testing"
)

var serverName = boot.RootPath() + "/MingServer"

func TestMain(m *testing.M) {
	_ = boot.Boot([]string{serverName}, serverName)
	m.Run()
}

func TestQueryTeachers(t *testing.T) {
	whereMaps := map[string]string{
		"nickname": "张三",
		"order":    "createtime ASC",
	}

	newRepo := teacher_repository.NewTeacherRepo
	total, teachers, err := newRepo.QueryTeachers(0, 10, whereMaps)
	assert.Nil(t, err)
	dumps.Dump(total)
	dumps.Dump(teachers)

	boot.Destroy()
}
