package trpc

import (
	"gitlab.upchinaproduct.com/taf/go/dcache"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"gitlab.upchinaproduct.com/upgo/utils/log"
)

var DCacheName2Ptr = make(map[string]*DCacheStruct)

type DCacheStruct struct {
	Module string
	Prx    *dcache.CacheProxy
}

// 初始化DCache模块
func DCacheInit(conf *conf.Conf, dcDomain string) error {
	log.Cache.Infof("init DCache start ......")
	dbConfList := conf.GetDomain(dcDomain)
	for key, dcName := range dbConfList {
		log.Cache.Infof("init DCache for key::%d dcName::%s start......", key, dcName)
		oneDCacheMap := conf.GetMap(dcDomain + "/" + dcName)

		dCache := new(DCacheStruct)
		dCache.Module = oneDCacheMap["module"]
		dCache.Prx = dcache.NewCacheProxy(oneDCacheMap["obj"])

		DCacheName2Ptr[dcName] = dCache
		log.Cache.Infof("init DCache for key::%d dcName::%s end......", key, dcName)
	}
	log.Cache.Infof("init DCache DCacheName2Ptr::", DCacheName2Ptr)
	log.Cache.Infof("init DCache success ......")
	return nil
}

// 获取DCache实例
func GetDCache(dcName string) *DCacheStruct {
	return DCacheName2Ptr[dcName]
}
