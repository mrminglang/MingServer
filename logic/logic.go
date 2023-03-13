package logic

import (
	"context"
	"encoding/json"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"server/logic/esrpc"
	"server/repositories/caches/ming_cache"
	"server/repositories/es/es_repository"
	"server/repositories/mysql/teacher_repository"
	"server/taf-protocol/MingApp"
	"server/utils/log"
	"strconv"
	"time"
)

// 初始化业务逻辑
func Init(conf *conf.Conf) {

	// ESDriverServer rpc
	esObj := conf.GetString("/obj/<esObj>")
	log.Es.Infof("esrpc.Init esObj::", esObj)
	//初始化es模块
	esrpc.Init(esObj)
}

// 获取老师列表
func GetTeacherList(_ context.Context, req *MingApp.GetTeacherListReq, rsp *MingApp.GetTeacherListRsp) (ret int32, err error) {
	startTime := time.Now().UnixMilli()
	log.Def.Infof("{GetTeacherList start req}|%s", req.Display())
	if req.BeginIndex < 0 || req.Count <= 0 {
		log.Def.Errorf("{GetTeacherList req param is failed}|%s|%d", "参数错误！", time.Now().UnixMilli()-startTime)
		return
	}

	// gorm 查询逻辑
	whereMaps := map[string]string{
		"order": "createtime ASC",
	}
	if req.NickName != "" {
		whereMaps["nickname"] = req.NickName
	}
	total, teachers, err := teacher_repository.QueryTeachers(int(req.BeginIndex), int(req.Count), whereMaps)
	if err != nil {
		log.Def.Errorf("{GetTeacherList QueryTeachers error}|%s|%d", err.Error(), time.Now().UnixMilli()-startTime)
		return ret, nil
	}
	if total <= 0 {
		log.Def.Errorf("{GetTeacherList QueryTeachers rsp is null}|%s|%d", "请求数据为空！", time.Now().UnixMilli()-startTime)
		return
	}

	rsp.Sum = int32(total)
	for _, teacher := range teachers {
		tmp := MingApp.TeacherList{
			NickName:   teacher.Nickname,
			HearderPic: teacher.Hearderpic,
		}
		rsp.Teachers = append(rsp.Teachers, tmp)
	}

	log.Def.Infof("{GetTeacherList end rsp}|%s|%d", rsp.Display(), time.Now().UnixMilli()-startTime)
	return ret, nil
}

// 设置DCache缓存
func SetStringCache(_ context.Context, req *MingApp.SetStringCacheReq, rsp *MingApp.SetStringCacheRsp) (ret int32, err error) {
	log.Cache.Infof("{SetStringCache start req}|%s", req.Display())
	rsp.Ret = ret
	if req.CacheKey == "" || req.CacheValue == "" {
		rsp.Msg = "参数错误"
		log.Cache.Errorf("{SetStringCache req param is failed}|%s", rsp.Msg)
		return
	}

	ret, err = ming_cache.SetStringCache(req.CacheKey, req.CacheValue)
	if err != nil {
		rsp.Msg = err.Error()
		log.Cache.Errorf("{SetStringCache is error}|%s", err.Error())
		return ret, nil
	}

	rsp.Msg = "success"
	log.Cache.Infof("{SetStringCache end rsp}|%s", rsp.Display())
	return
}

// 获取DCache缓存
func GetStringCache(_ context.Context, req *MingApp.GetStringCacheReq, rsp *MingApp.GetStringCacheRsp) (ret int32, err error) {
	log.Cache.Infof("{GetStringCache req}|%s", req.Display())
	rsp.Ret = ret
	if req.CacheKey == "" {
		rsp.Msg = "参数错误"
		log.Cache.Errorf("{GetStringCache req param is failed}|%s", rsp.Msg)
		return
	}

	_, err, cacheRsp := ming_cache.GetStringCache(req.CacheKey)
	if err != nil {
		rsp.Msg = err.Error()
		log.Cache.Errorf("{GetStringCache is error}|%s", err.Error())
		return ret, nil
	}

	rsp.Msg = "success"
	rsp.CacheValue = cacheRsp.Value
	log.Cache.Infof("{GetStringCache end rsp}|%s", rsp.Display())
	return
}

// 设置ES数据
func SetESData(_ context.Context, req *MingApp.SetESDataReq, rsp *MingApp.SetESDataRsp) (ret int32, err error) {
	log.Es.Infof("{SetESData start req}|%s", req.Display())
	rsp.Ret = ret
	if req.IndexName == "" || req.Typ == "" || req.Id <= 0 {
		rsp.Msg = "参数错误"
		log.Es.Errorf("{SetESData req param is failed}|%s", rsp.Msg)
		return
	}
	ret, err = es_repository.SetESData(req.IndexName, req.Typ, strconv.Itoa(int(req.Id)), req.Teachers)
	if err != nil {
		rsp.Msg = err.Error()
		log.Es.Errorf("{SetESData is error}|%s", err.Error())
		return ret, nil
	}

	rsp.Msg = "success"
	log.Es.Infof("{SetESData end rsp}|%s", rsp.Display())
	return
}

// 获取ES数据 by id
func GetESDataById(_ context.Context, req *MingApp.GetESDataByIdReq, rsp *MingApp.GetESDataByIdRsp) (ret int32, err error) {
	log.Es.Infof("{GetESDataById start req}|%s", req.Display())
	rsp.Ret = ret
	if req.IndexName == "" || req.Typ == "" || req.Id <= 0 {
		rsp.Msg = "参数错误"
		log.Es.Errorf("{GetESDataById req param is failed}|%s", rsp.Msg)
		return
	}

	source, err := es_repository.GetESDataById(req.IndexName, req.Typ, strconv.Itoa(int(req.Id)))
	if err != nil {
		rsp.Msg = err.Error()
		log.Es.Errorf("{GetESDataById is error}|%s", err.Error())
		return ret, nil
	}
	err = json.Unmarshal([]byte(source), &rsp.Teachers)
	if err != nil {
		log.Es.Errorf("{GetESDataById source Unmarshal error::%s}", err.Error())
		return 0, nil
	}

	rsp.Msg = "success"
	log.Es.Infof("{GetESDataById end rsp}|%s", rsp.Display())
	return
}
