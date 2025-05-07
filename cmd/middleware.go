package cmd

import (
	"ewallet-ums/helpers"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {

	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("Authorization header is missing")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Authorization header is missing", nil)
		return
	}

	fmt.Println("Authorization header:", auth)

	_, err := d.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}
	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized validate token", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("Token expired", claim.ExpiresAt)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Token expired", nil)
		return
	}

	ctx.Set("token", claim)

	ctx.Next()
	return
}

func (d *Dependency) MiddlewareValidateRefreshToken(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("Authorization header is empty")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return

	}

	_, err := d.UserRepository.GetUserSessionByRefreshToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized validate token", nil)
		ctx.Abort()
		return

	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("Token expired", claim.ExpiresAt)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Token expired", nil)
		ctx.Abort()
		return
	}
	ctx.Set("token", claim)
	ctx.Next()
}
