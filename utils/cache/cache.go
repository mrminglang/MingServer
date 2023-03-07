package cache

import (
	"MingServer/utils/log"
	"gitlab.upchinaproduct.com/taf/go/dcache"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
)

// Ming缓存dcache配置
var MingCacheModule = ""
var MingCachePrx *dcache.CacheProxy

// 初始化DCache模块
func Init(conf *conf.Conf) {
	log.Cache.Infof("init DCache start ......")
	// Ming缓存dcache初始化
	MingCacheModule = conf.GetString("/app/<mingCacheModule>")
	MingCachePrx = dcache.NewCacheProxy(conf.GetString("/obj/<mingCacheObj>"))
	log.Cache.Infof("init DCache success ......")
}
