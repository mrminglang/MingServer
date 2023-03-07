package teacher_repository_test

import (
	"github.com/mrminglang/tools/dumps"
	"testing"
)

func TestMain(m *testing.M) {
	//boot.Boot()
	m.Run()
}

// go test -v -run TestQueryTeachers teacher_test.go
func TestQueryTeachers(t *testing.T) {
	whereMaps := map[string]string{
		"nickname": "张三",
		"order":    "createtime ASC",
	}

	dumps.Dump(whereMaps)

	//total, teachers, err := teacher_repository.NewTeacher().QueryTeachers(0, 10, whereMaps)
	//assert.Nil(t, err)
	//dumps.Dump(total)
	//dumps.Dump(teachers)

	//boot.Destroy()
}
