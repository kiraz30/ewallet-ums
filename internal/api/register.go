package api

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisterService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Info("Failded to parse requst:", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Info("Failded to validate:", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	resp, err := api.RegisterService.Regiter(c.Request.Context(), req)
	if err != nil {
		log.Error("Failed to register user:", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
	return
}
