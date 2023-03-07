package models

//投顾老师信息db结构
type Teacher struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Nickname   string `gorm:"column:nickname"`   // 昵称
	Hearderpic string `gorm:"column:hearderpic"` // 头像
	Tgupname   string `gorm:"column:tgupname"`   // 关联投顾账号
	Issync     int    `gorm:"column:issync"`     // 是否同步投顾内容
	Gender     int    `gorm:"column:gender"`     // 性别| 0：未知、1：男、2：女
	Desc       string `gorm:"column:desc"`       // 简介
	Uid        string `gorm:"column:uid"`        // 优品id
	Status     int    `gorm:"column:status"`     // 状态
	Createuser string `gorm:"column:createuser"` // 创建用户
	Createtime int    `gorm:"column:createtime"` // 创建时间
	Updateuser string `gorm:"column:updateuser"` // 修改人
	Updatetime int    `gorm:"column:updatetime"` // 修改时间
}

func (r *Teacher) TableName() string {
	return "t_teacher"
}
