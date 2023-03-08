package ormdb

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//初始化完成标记
var IsInitOk = false

//des解密key
// var key = []byte("0My@Sql1")

//数据库名->数据库指针
var DbConnName2Ptr = make(map[string]*gorm.DB)

//初始化配置文件里面的数据库到map中
func Init(conf *conf.Conf, dbDoman string) (err error) {
	dbConfList := conf.GetDomain(dbDoman)
	for _, dbName := range dbConfList {
		oneDbConfMap := conf.GetMap(dbDoman + "/" + dbName)
		switch oneDbConfMap["dbType"] {
		case "mysql":
			{
				oneDbConStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
					oneDbConfMap["user"], oneDbConfMap["password"], oneDbConfMap["host"],
					oneDbConfMap["port"], oneDbConfMap["db_name"])

				fmt.Printf("{startInitDb dbType^dbName^dbConfStr}|mysql|%s|%s\n", dbName, oneDbConStr)
				dbPtr, Err := gorm.Open(mysql.Open(oneDbConStr), &gorm.Config{
					Logger:                                   logger.Default.LogMode(logger.Info),
					DisableForeignKeyConstraintWhenMigrating: true,
					NamingStrategy: schema.NamingStrategy{
						TablePrefix:   "public.",
						SingularTable: true, //表名后面不加s
					},
				})
				if Err != nil {
					fmt.Printf("{initDbError dbName^Err}|%s|%s\n", dbName, Err.Error())
					continue
				} else {
					//连接池配置
					//这个sqlDb和dbPtr这个不一样 前者是sql.DB 后者 是gorm.DB sql.DB为通用数据库接口
					sqlDb, _ := dbPtr.DB()
					sqlDb.SetMaxIdleConns(30)  //设置空闲连接池中连接的最大数量
					sqlDb.SetMaxOpenConns(100) //设置打开数据库连接的最大数量

					DbConnName2Ptr[dbName] = dbPtr
					fmt.Printf("{finishInitDb dbName}|%s\n", dbName)
				}
			}
		}
		fmt.Println(DbConnName2Ptr)
	}
	fmt.Println("{finishAllDbInit}")

	IsInitOk = true
	return nil
}

func GetDb(dbConnName string) *gorm.DB {
	return DbConnName2Ptr[dbConnName]
}
