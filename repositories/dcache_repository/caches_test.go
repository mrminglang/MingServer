package dcache_repository_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/repositories/dcache_repository"
	"testing"
)

var serverName = boot.RootPath() + "/MingServer"

func TestMain(m *testing.M) {
	_ = boot.Boot([]string{serverName}, serverName)
	m.Run()
}

func TestSetStringCache(t *testing.T) {
	key := "key-20230419"
	value := "value-20230419"

	newRepo := dcache_repository.NewDCacheRepo
	ret, err := newRepo.SetStringCache(key, value, "ming")
	if err != nil {
		assert.Error(t, err)
		return
	}

	dumps.Dump(ret)
}

func TestGetStringCache(t *testing.T) {
	newRepo := dcache_repository.NewDCacheRepo

	key := "key-20230419"
	ret, err, rsp := newRepo.GetStringCache(key, "ming")
	if err != nil {
		assert.Error(t, err)
		return
	}
	dumps.Dump(ret)
	dumps.Dump(rsp)

	key1 := "NewsIdGen"
	ret1, err, rsp1 := newRepo.GetStringCache(key1, "cnews")
	if err != nil {
		assert.Error(t, err)
		return
	}

	dumps.Dump(ret1)
	dumps.Dump(rsp1)

}
