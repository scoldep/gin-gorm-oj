// package middlewares
//
// import (
//
//	"gin-gorm-oj/helper"
//	"github.com/gin-gonic/gin"
//	"net/http"
//
// )
//
//	func AuthUserCheck() gin.HandlerFunc {
//		return func(c *gin.Context) {
//			// TODD: Check if user is admin
//			auth := c.GetHeader("Authorization")
//			userClaim, err := helper.AnalyseToken(auth)
//			if err != nil {
//				c.Abort()
//				c.JSON(http.StatusOK, gin.H{
//					"code": http.StatusUnauthorized,
//					"msg":  "Unauthorized, Authorization",
//				})
//				return
//			}
//			if userClaim == nil {
//				c.Abort()
//				c.JSON(http.StatusOK, gin.H{
//					"code": http.StatusUnauthorized,
//					"msg":  "Unauthorized Admin",
//				})
//				return
//			}
//			c.Set("user_claims", userClaim)
//			c.Next()
//		}
//	}
package middlewares

import (
	"gin-gorm-oj/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthUserCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			return
		}
		if userClaim == nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Admin",
			})
			return
		}
		c.Set("user_claims", userClaim)
		c.Next()
	}
}
