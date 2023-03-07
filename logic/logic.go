package logic

import (
	"MingServer/repositories/caches/ming_cache"
	"MingServer/repositories/mysql/teacher_repository"
	"MingServer/taf-protocol/MingApp"
	"MingServer/utils/log"
	"context"
	"time"
)

// 获取老师列表
func GetTeacherList(_ context.Context, req *MingApp.GetTeacherListReq, rsp *MingApp.GetTeacherListRsp) (ret int32, err error) {
	startTime := time.Now().UnixMilli()
	log.Def.Infof("{GetTeacherList start req}|%s", req.Display())
	if req.BeginIndex < 0 || req.Count <= 0 {
		log.Def.Errorf("{GetTeacherList req param is failed}|%s|%d", "参数错误！", time.Now().UnixMilli()-startTime)
		return
	}

	// 设置缓存
	key := "mingming1000"
	value := "mingmingmingming"
	ret, err = ming_cache.SetStringCache(key, value)
	if err != nil {
		log.Cache.Errorf("{GetTeacherList SetStringCache is error}|%s|%d", err.Error(), time.Now().UnixMilli()-startTime)
		return 0, nil
	}
	// 获取缓存
	ret, err, cacheRsp := ming_cache.GetStringCache(key)
	if err != nil {
		return 0, nil
	}
	rsp.CacheValue = cacheRsp.Value
	log.Cache.Infof("{GetTeacherList GetStringCache cacheRsp::}|%s", cacheRsp)

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
