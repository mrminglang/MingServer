package trpc

import (
	"gitlab.upchinaproduct.com/taf/go/dcache"
	"server/utils/log"
)

// Ming缓存dcache配置
var MCModule = ""
var MCPrx *dcache.CacheProxy

// 初始化DCache模块
func MCInit(module string, obj string) error {
	log.Cache.Infof("init DCache start ......")
	// Ming缓存dcache初始化
	MCModule = module
	MCPrx = dcache.NewCacheProxy(obj)
	log.Cache.Infof("init DCache success ......")

	return nil
}
