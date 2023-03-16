package log

import (
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
)

// 预定义滚动日志（默认支持日志级别）框架全局滚动日志：taf.LOG()
var Roll = taf.GetLogger("roll")

// 预定义按天日志（默认不支持日志级别，需要设置参数）
var Def = taf.GetDayLogger("def", 1)     // 默认日志
var Conf = taf.GetDayLogger("conf", 1)   // 配置日志
var Cache = taf.GetDayLogger("cache", 1) // 缓存类日志
var Es = taf.GetDayLogger("es", 1)
var Data = taf.GetDayLogger("data", 1)

// 启动日志设置
func Init() {
	Def.SetShowLevel(true)  // 设置日志级别
	Def.SetCallerFlag(true) // 显示代码位置

	Conf.SetShowLevel(true)
	Conf.SetCallerFlag(true)

	Cache.SetShowLevel(true)
	Cache.SetCallerFlag(true)

	Es.SetShowLevel(true)
	Es.SetCallerFlag(true)

	Data.SetShowLevel(true)
	Data.SetCallerFlag(true)
}

// 关闭日志设置
func Destroy() {

}
