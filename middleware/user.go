package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo_list/global"
	"todo_list/initjwt"
	"todo_list/model"
)

func UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "Without Authorization",
				"detail":  "Authorization Is Empty",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "Authorization Format Error",
				"detail":  "Without Bearer Header",
			})
			c.Abort()
			return
		}

		// 提取token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析JWT token
		claims, err := initjwt.PraseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   -1,
				"msg":    "Error Occurred When PraseToken",
				"detail": err.Error(),
			})
			c.Abort()
			return
		}

		// 查询用户信息
		var user model.User
		querySql := "SELECT UserID, UserName FROM user WHERE UserID = ?"
		err = global.Db.QueryRow(querySql, claims.UserID).Scan(&user.UserID, &user.UserName)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":   -1,
					"msg":    "The user does not exist.",
					"detail": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":   -1,
				"msg":    "Query failed",
				"detail": err.Error(),
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user", user)
		c.Set("user_id", user.UserID)
		c.Next()
	}
}
