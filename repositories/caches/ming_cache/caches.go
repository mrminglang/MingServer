package ming_cache

import (
	"gitlab.upchinaproduct.com/taf/go/dcache"
	"server/utils/log"
	"server/utils/trpc"
)

type User struct {
	Name string
}

// 设置缓存 string: key -> value
func SetStringCache(key string, value string) (ret int32, err error) {
	req := dcache.SetStringReq{Key: key, Value: value}
	log.Cache.Infof("{SetStringCache start req::}", req)
	ret, err = trpc.MCPrx.SetString(trpc.MCModule, &req)
	if err != nil {
		log.Cache.Errorf("{SetStringCache failed error::}|%s", err.Error())
		return
	}
	log.Cache.Infof("{SetStringCache end ret::}", ret)
	return
}

// 获取缓存 string: key -> vaule
func GetStringCache(key string) (ret int32, err error, rsp dcache.GetStringRsp) {
	req := dcache.GetReq{Key: key}
	log.Cache.Infof("{GetStringCache start req::}", req)

	ret, err = trpc.MCPrx.GetString(trpc.MCModule, &req, &rsp)
	if err != nil {
		log.Cache.Errorf("{GetStringCache failed error::}|%s", err.Error())
		return
	}
	log.Cache.Infof("{GetStringCache end rsp::}", rsp)
	return
}

// 设置缓存 StructEx: key -> value
func SetStructExCache(key string, value User) (ret int32, err error) {
	req := dcache.SetStructReq{
		Key: key,
		//Value: &value,
	}

	log.Cache.Infof("{SetStructCache start key::%s, value::%v}", key, value)
	ret, err = trpc.MCPrx.SetStructEx(trpc.MCModule, &req)
	if err != nil {
		log.Cache.Errorf("{SetStructCache failed error::}|%s", err.Error())
		return
	}
	log.Cache.Infof("{SetStructCache end ret::}|%d", ret)
	return
}
