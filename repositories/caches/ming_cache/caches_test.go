package ming_cache_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/repositories/caches/ming_cache"
	"testing"
)

var serverName = boot.RootPath() + "/MingServer"

func TestMain(m *testing.M) {
	_ = boot.Boot([]string{serverName}, serverName)
	m.Run()
}

func TestSetStringCache(t *testing.T) {
	key := "key-20230313"
	value := "value-20230313"

	ret, err := ming_cache.SetStringCache(key, value)
	if err != nil {
		assert.Error(t, err)
		return
	}

	dumps.Dump(ret)
}

func TestGetStringCache(t *testing.T) {
	key := "key-20230313"
	ret, err, rsp := ming_cache.GetStringCache(key)
	if err != nil {
		assert.Error(t, err)
		return
	}

	dumps.Dump(ret)
	dumps.Dump(rsp)
}
