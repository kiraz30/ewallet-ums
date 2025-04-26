package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("authorization")
	if auth == "" {
		log.Println("Authorization empty")
		helpers.SendResponseHTTP(ctx, 401, "Unauthorized", nil)
		return
	}

	if d.UserRepo == nil {
		log.Println("UserRepo is nil")
		helpers.SendResponseHTTP(ctx, 500, "Internal Server Error", nil)
		ctx.Abort()
		return
	}

	_, err := d.UserRepo.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, 401, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx, auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, 401, "Unauthorized", nil)
		ctx.Abort()
		return
	}
	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("Token expired")
		helpers.SendResponseHTTP(ctx, 401, "Unauthorized", nil)
		return
	}

	ctx.Set("token", claim)
	ctx.Next()
}
