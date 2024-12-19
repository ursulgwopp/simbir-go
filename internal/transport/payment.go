package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

// @Router /api/Payment/Hesoyam/{id} [post]
// @Security ApiKeyAuth
// @Summary Hesoyam
// @Tags Hesoyam
// @Description Deposit
// @ID hesoyam
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) hesoyam(c *gin.Context) {
	accountId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	isAdmin, err := getIsAdmin(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.Hesoyam(accountId, userId, isAdmin); err != nil {
		if errors.Is(err, custom_errors.ErrAccessDenied) {
			models.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "successfully"})
}
