package api

import (
	"ewallet-ums/constans"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	refreshToken := c.Request.Header.Get("Authorization")
	claim, ok := c.Get("token")
	if !ok {
		log.Info("Failded to get claim in context")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constans.ErrServerError, nil)
		return
	}

	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Info("Failded to parse claim to claimToken")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constans.ErrServerError, nil)
		return
	}
	resp, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Info("Failded on Refresh Token service:", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constans.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constans.SuccessMessage, resp)
}
