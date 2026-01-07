package logic

import (
	"context"
	"encoding/json"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"gitlab.upchinaproduct.com/upgo/utils/esdb/es_repository"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/log"
	"gitlab.upchinaproduct.com/upgo/utils/tafrpc/dcache_rpc"
	"server/repositories/teacher_repository"
	"server/taf-protocol/MingApp"
	"strconv"
	"time"
)

// 初始化业务逻辑
func Init(conf *conf.Conf) error {
	log.Data.Infof("{logic init start......}")

	// 初始化数据仓库
	teacher_repository.InitTeacherRepo(conf)
	dcache_rpc.InitDCacheRepo(conf)

	log.Data.Infof("{logic init success......}")

	return nil
}

// 获取老师列表
func GetTeacherList(_ context.Context, req *MingApp.GetTeacherListReq, rsp *MingApp.GetTeacherListRsp) (ret int32, err error) {
	startTime := time.Now().UnixMilli()
	log.Data.Infof("{GetTeacherList start req}|%s", req.Display())
	if req.BeginIndex < 0 || req.Count <= 0 {
		log.Data.Errorf("{GetTeacherList req param is failed}|%s|%d", "参数错误！", time.Now().UnixMilli()-startTime)
		return
	}

	// gorm 查询逻辑
	whereMaps := map[string]string{
		"order": "createtime ASC",
	}
	if req.NickName != "" {
		whereMaps["nickname"] = req.NickName
	}
	newRepo := teacher_repository.NewTeacherRepo
	total, teachers, err := newRepo.QueryTeachers(int(req.BeginIndex), int(req.Count), whereMaps)
	if err != nil {
		log.Data.Errorf("{GetTeacherList QueryTeachers error}|%s|%d", err.Error(), time.Now().UnixMilli()-startTime)
		return ret, nil
	}
	if total <= 0 {
		log.Data.Errorf("{GetTeacherList QueryTeachers rsp is null}|%s|%d", "请求数据为空！", time.Now().UnixMilli()-startTime)
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

	log.Data.Infof("{GetTeacherList end rsp}|%s|%d", rsp.Display(), time.Now().UnixMilli()-startTime)
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
	newRepo := dcache_rpc.NewDCacheRepo
	ret, err = newRepo.SetString(req.CacheKey, req.CacheValue, "ming")
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
	newRepo := dcache_rpc.NewDCacheRepo
	_, err, cacheRsp := newRepo.GetString(req.CacheKey, "ming")
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
		return ret, nil
	}

	// 先判断索引是否存在，不存在则创建索引
	if ok, _ := es_repository.IndexExists([]string{req.IndexName}); !ok {
		index, err := es_repository.CreateIndex(req.IndexName, "")
		if err != nil {
			log.Es.Errorf("{SetESData CreateIndex error::}", err.Error())
			return ret, nil
		}
		log.Es.Infof("{SetESData CreateIndex index::}", index)
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
