package teacher_repository

import (
	"MingServer/repositories/models"
	"MingServer/utils/ormdb"
)

// QueryTeachers 分页查询老师
func QueryTeachers(
	beginIndex int,
	count int,
	whereMaps map[string]string,
) (total int64, teachers []*models.Teacher, err error) {
	query := ormdb.GetDb("daka").Model(&teachers)

	// 昵称
	if whereMaps["nickname"] != "" {
		query = query.Where("nickname LIKE ?", "%"+whereMaps["nickname"]+"%")
	}
	// 排序
	if whereMaps["order"] != "" {
		query = query.Order(whereMaps["order"])
	}
	// 分页查询
	err = query.Count(&total).Offset(beginIndex).Limit(count).Find(&teachers).Error
	if err != nil {
		return
	}

	return
}
