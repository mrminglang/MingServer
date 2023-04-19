package teacher_repository

import (
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"gorm.io/gorm"
	"server/repositories/models"
	"server/utils/log"
	"server/utils/ormdb"
)

var NewTeacherRepo *NewTeacherRepoStruct

type NewTeacherRepoStruct struct {
	DbConn *gorm.DB
}

// 初始化数据仓库
func InitTeacherRepo(conf *conf.Conf) {
	log.Db.Infof("{teacher_repository InitTeacherRepo start......}")

	newRepo := new(NewTeacherRepoStruct)
	newRepo.DbConn = ormdb.GetDb("daka")
	NewTeacherRepo = newRepo

	log.Db.Infof("{teacher_repository InitTeacherRepo end......}")
}

// QueryTeachers 分页查询老师
func (s *NewTeacherRepoStruct) QueryTeachers(
	beginIndex int,
	count int,
	whereMaps map[string]string,
) (total int64, teachers []*models.Teacher, err error) {
	query := s.DbConn.Model(&teachers)

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
