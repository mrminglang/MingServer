package models

import (
	"encoding/json"
	"net/http"
)

// AuthRsp 认证返回结构体
type AuthRsp struct {
	Code      int32  `json:"code"`       // 1 成功，<0 失败
	Note      string `json:"note"`       // 备注
	Token     string `json:"token"`      // 获取到的凭证，code=1时才会返回token
	ExpiresIn int32  `json:"expires_in"` // 凭证有效时间，单位：秒
}

// HTTPResponseHandler 用于处理HTTP响应的通用结构
type HTTPResponseHandler struct {
	writer   http.ResponseWriter
	encoder  *json.Encoder
	funcName string
}

// NewHTTPResponseHandler 创建一个新的HTTP响应处理器
func NewHTTPResponseHandler(w http.ResponseWriter, funcName string) *HTTPResponseHandler {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return &HTTPResponseHandler{
		writer:   w,
		encoder:  json.NewEncoder(w),
		funcName: funcName,
	}
}

// WriteResponse 写入响应数据
func (h *HTTPResponseHandler) WriteResponse(response interface{}) error {
	return h.encoder.Encode(response)
}

// GetEncoder 获取JSON编码器
func (h *HTTPResponseHandler) GetEncoder() *json.Encoder {
	return h.encoder
}

// GetFuncName 获取函数名称
func (h *HTTPResponseHandler) GetFuncName() string {
	return h.funcName
}
