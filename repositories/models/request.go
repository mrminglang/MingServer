package models

// AuthReq 认证请求结构体
type AuthReq struct {
	LoginId string `json:"loginId"` // 用户登录ID（消息源标识）
	Sign    string `json:"sign"`    // 请求参数进行签名，签名方式见下面的说明
}
