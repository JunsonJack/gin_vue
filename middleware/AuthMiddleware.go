package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"junsonjack.cn/go_vue/common"
	"junsonjack.cn/go_vue/model"
)

func AuthMiddleware() gin.HandlerFunc{
	return func (c *gin.Context)  {
		// 获取authorization header 
		tokenString := c.GetHeader("Authorization")

		// 校验token 
		// oauth2.0 规定Authorization的字符串开头必须要有Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			c.JSON(http.StatusUnauthorized,gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			c.Abort() //抛弃本次请求
			return
		}
		tokenString = tokenString[7:] //Bearer+空格一共占了七位

		token , claims , err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized,gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			c.Abort() //抛弃本次请求
			return
		}

		// 验证通过以后获取claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)

		// 用户不存在
		if user.ID == 0{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			c.Abort() //抛弃本次请求
			return
		}
		// 用户存在，将user的信息写入上下文
		c.Set("user",user)

		c.Next()


	}
}