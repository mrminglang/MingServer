package ming_cache_test

import (
	"MingServer/repositories/caches/ming_cache"
	"github.com/mrminglang/tools/dumps"
	genid "github.com/srlemon/gen-id"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetStructExCache(t *testing.T) {

	key := "20230303"
	value := ming_cache.User{
		Name: genid.NewGeneratorData().Name,
	}

	ret, err := ming_cache.SetStructExCache(key, value)
	if err != nil {
		assert.Error(t, err)
		return
	}

	dumps.Dump(ret)
}
