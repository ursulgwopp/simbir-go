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
// @Summary AdminListAccounts
// @Tags Admin Account
// @Description List Accounts
// @ID admin-list-accounts
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
		if errors.Is(err, custom_errors.ErrInvalidPaginationParams) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// @Router /api/Admin/Account/{id} [get]
// @Security ApiKeyAuth
// @Summary AdminGetAccount
// @Tags Admin Account
// @Description Get Account
// @ID admin-get-account
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
		if errors.Is(err, custom_errors.ErrAccountIdNotFound) {
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
// @Summary AdminCreateAccount
// @Tags Admin Account
// @Description Create Account
// @ID admin-create-account
// @Accept json
// @Produce json
// @Param Input body models.AdminAccountRequest true "Account Info"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminCreateAccount(c *gin.Context) {
	var req models.AdminAccountRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.AdminCreateAccount(req)
	if err != nil {
		if errors.Is(err, custom_errors.ErrInvalidUsernameLength) ||
			errors.Is(err, custom_errors.ErrInvalidUsernameCharacters) ||
			errors.Is(err, custom_errors.ErrInvalidPasswordLength) ||
			errors.Is(err, custom_errors.ErrInvalidBalanceValue) ||
			errors.Is(err, custom_errors.ErrUsernameIsNotUnique) {
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
// @Summary AdminUpdateAccount
// @Tags Admin Account
// @Description Update Account
// @ID admin-update-account
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
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.AdminUpdateAccount(accountId, req); err != nil {
		if errors.Is(err, custom_errors.ErrAccountIdNotFound) ||
			errors.Is(err, custom_errors.ErrInvalidUsernameLength) ||
			errors.Is(err, custom_errors.ErrInvalidUsernameCharacters) ||
			errors.Is(err, custom_errors.ErrInvalidPasswordLength) ||
			errors.Is(err, custom_errors.ErrUsernameIsNotUnique) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "account successfully updated"})
}

// @Router /api/Admin/Account/{id} [delete]
// @Security ApiKeyAuth
// @Summary AdminDeleteAccount
// @Tags Admin Account
// @Description Delete Account
// @ID admin-delete-account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminDeleteAccount(c *gin.Context) {
	userId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accountId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//////////////////////////////////////////////////////////////////////////////////////////
	if accountId == userId {
		models.NewErrorResponse(c, http.StatusBadRequest, "can not delete admin account")
		return
	}
	//////////////////////////////////////////////////////////////////////////////////////////

	if err := t.service.AdminDeleteAccount(accountId); err != nil {
		if errors.Is(err, custom_errors.ErrAccountIdNotFound) ||
			errors.Is(err, custom_errors.ErrCanNotDeleteAccount) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "account successfully deleted"})
}
