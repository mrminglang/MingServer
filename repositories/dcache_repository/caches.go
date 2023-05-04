package dcache_repository

import (
	"gitlab.upchinaproduct.com/taf/go/dcache"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/log"
	"server/utils/trpc"
)

var NewDCacheRepo *NewDCacheStruct

type NewDCacheStruct struct {
	DCache map[string]*trpc.DCacheStruct
}

// 初始化DCache数据仓库
func InitDCacheRepo(conf *conf.Conf) {
	log.Cache.Infof("{dcache_repository InitDCacheRepo start......}")
	newRepo := new(NewDCacheStruct)
	newRepo.DCache = trpc.DCacheName2Ptr
	NewDCacheRepo = newRepo
	log.Cache.Infof("{dcache_repository InitDCacheRepo end......}")
}

// 设置缓存 string: key -> value
func (s *NewDCacheStruct) SetStringCache(key string, value string, name string) (ret int32, err error) {
	req := dcache.SetStringReq{Key: key, Value: value}
	log.Cache.Infof("{SetStringCache start req::}", req)
	ret, err = s.DCache[name].Prx.SetString(s.DCache[name].Module, &req)
	if err != nil {
		log.Cache.Errorf("{SetStringCache failed error::}|%s", err.Error())
		return
	}
	log.Cache.Infof("{SetStringCache end ret::}", ret)
	return
}

// 获取缓存 string: key -> vaule
func (s *NewDCacheStruct) GetStringCache(key string, name string) (ret int32, err error, rsp dcache.GetStringRsp) {
	req := dcache.GetReq{Key: key}
	log.Cache.Infof("{GetStringCache start req::}", req)

	ret, err = s.DCache[name].Prx.GetString(s.DCache[name].Module, &req, &rsp)
	if err != nil {
		log.Cache.Errorf("{GetStringCache failed error::}|%s", err.Error())
		return
	}
	log.Cache.Infof("{GetStringCache end rsp::}", rsp)
	return
}
