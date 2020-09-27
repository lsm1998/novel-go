package logic

import (
	"api-gateway/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
	"utils"
)

// 最多每秒允许 limit 个请求
var lmt = rate.NewLimiter(rate.Limit(config.Config.Gateway.Limit), 1)

// 鉴权
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		uidStr := c.GetHeader("uid")  // 用户uid
		token := c.GetHeader("token") // 访问令牌
		var errMsg string
		var loginClaims *utils.LoginClaims
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			errMsg = "请求未包含有效用户ID"
			goto FAIT
		}
		loginClaims, err = utils.ValidToken(token, uid)
		if err != nil {
			errMsg = err.Error()
			goto FAIT
		}
		// 验证通过，会继续访问下一个中间件
		c.Set("roles", loginClaims.Role)
		c.Next()
		return
	FAIT:
		// 验证不通过，不再调用后续的函数处理
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": errMsg})
	}
}

// 限流
func LimitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !lmt.AllowN(time.Now(), 1) {
			c.JSON(200, gin.H{
				"code":    utils.ERR_LIMIT,
				"message": "服务端繁忙",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func AccessRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		println(c.Request.URL.Path)
		c.ClientIP()
	}
}
