package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageResult struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

// Success 请求成功返回
func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

func AuthFailed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{401, message, 0})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}

// SuccessPage 请求成功返回分页结果
func SuccessPage(message string, data interface{}, total int, c *gin.Context) {
	page := &PageResult{Total: total, List: data}
	c.JSON(http.StatusOK, Response{200, message, page})
}
