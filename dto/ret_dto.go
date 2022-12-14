package dto

import "github.com/gin-gonic/gin"

// RetStruct response ret struct refactoring
type RetStruct struct {
	Ret  bool   `json:"ret"`
	Data gin.H  `json:"data"` //数据
	Code uint   `json:"code"` //状态码
	Msg  string `json:"msg"`  //信息
}
