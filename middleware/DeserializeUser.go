package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/AhmedSaIah/fiber/config"
	"github.com/AhmedSaIah/fiber/services"
	"github.com/AhmedSaIah/fiber/utils"
)

func DeserializeUser(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		cnf, err := config.LoadConfig(".")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		sub, err := utils.ValidateToken(accessToken, cnf.AccessTokenPublicKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		}

		user, err := userService.FindUserById(fmt.Sprint(sub))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no longer exists"})
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
