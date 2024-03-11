package gtoken

import (
  "github.com/gogf/gf/v2/net/ghttp"
  "github.com/gogf/gf/v2/text/gstr"
)

// TokenConfig 登录令牌配置
type TokenConfig struct {
  SecretKey       string `json:"secretKey"`
  Expires         int64  `json:"expires"`
  AutoRefresh     bool   `json:"autoRefresh"`
  RefreshInterval int64  `json:"refreshInterval"`
  MaxRefreshTimes int64  `json:"maxRefreshTimes"`
  MultiLogin      bool   `json:"multiLogin"`
}

// GetAuthorization 获取authorization
func GetAuthorization(r *ghttp.Request) string {
  // 默认从请求头获取
  var authorization = r.Header.Get("Authorization")

  // 如果请求头不存在则从get参数获取
  if authorization == "" {
    return r.Get("authorization").String()
  }
  return gstr.Replace(authorization, "Bearer ", "")
}
