package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

// @Router /api/Admin/Account [get]
// @Security ApiKeyAuth
// @Summary ListAccounts
// @Tags Admin Account
// @Description List Accounts
// @ID list-accounts
// @Accept json
// @Produce json
// @Param from query int true "From"
// @Param count query int true "Count"
// @Success 200 {array} models.AdminAccountResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminListAccounts(c *gin.Context) {
	from_ := c.Query("from")
	count_ := c.Query("count")

	from, err := strconv.Atoi(from_)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	count, err := strconv.Atoi(count_)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accounts, err := t.service.AdminListAccounts(from, count)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// @Router /api/Admin/Account/{id} [get]
// @Security ApiKeyAuth
// @Summary GetAccount
// @Tags Admin Account
// @Description Get Account
// @ID get-account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.AdminAccountResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminGetAccount(c *gin.Context) {
	accountId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := t.service.AdminGetAccount(accountId)
	if err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Router /api/Admin/Account [post]
// @Security ApiKeyAuth
// @Summary CreateAccount
// @Tags Admin Account
// @Description Create Account
// @ID create-account
// @Accept json
// @Produce json
// @Param Input body models.AdminAccountRequest true "Account Info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminCreateAccount(c *gin.Context) {
	var req models.AdminAccountRequest
	if err := c.BindJSON(req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.AdminCreateAccount(req)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUsernameExists) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
}

// @Router /api/Admin/Account/{id} [put]
// @Security ApiKeyAuth
// @Summary UpdateAccount
// @Tags Admin Account
// @Description Update Account
// @ID update-account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param Input body models.AdminAccountRequest true "Update Info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminUpdateAccount(c *gin.Context) {
	accountId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req models.AdminAccountRequest
	if err := c.BindJSON(req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.AdminUpdateAccount(accountId, req); err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) || errors.Is(err, custom_errors.ErrUsernameExists) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "account successfully updated")
}

// @Router /api/Admin/Account/{id} [delete]
// @Security ApiKeyAuth
// @Summary DeleteAccount
// @Tags Admin Account
// @Description Delete Account
// @ID delete-account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminDeleteAccount(c *gin.Context) {
	accountId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.AdminDeleteAccount(accountId); err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
