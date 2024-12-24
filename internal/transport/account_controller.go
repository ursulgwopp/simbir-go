package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

// @Router /api/Account/SignUp [post]
// @Summary SignUp
// @Tags Account
// @Description Sign Up
// @ID sign-up
// @Accept json
// @Produce json
// @Param Input body models.AccountRequest true "SignUp Info"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) signUp(c *gin.Context) {
	var req models.AccountRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.SignUp(req)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUsernameExists) || errors.Is(err, custom_errors.ErrInvalidParams) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
}

// @Router /api/Account/SignIn [post]
// @Summary SignIn
// @Tags Account
// @Description Sign In
// @ID sign-in
// @Accept json
// @Produce json
// @Param Input body models.AccountRequest true "SignIn Info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) signIn(c *gin.Context) {
	var req models.AccountRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := t.service.SignIn(req)
	if err != nil {
		if errors.Is(err, custom_errors.ErrInvalidUsernameOrPassword) {
			models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrInvalidUsernameOrPassword.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

// @Router /api/Account/SignOut [post]
// @Security ApiKeyAuth
// @Summary SignOut
// @Tags Account
// @Description Sign Out
// @ID sign-out
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) signOut(c *gin.Context) {
	token, err := getToken(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.SignOut(token); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "successfully signed out"})
}

// @Router /api/Account/Me [get]
// @Security ApiKeyAuth
// @Summary GetProfile
// @Tags Account
// @Description Get Profile
// @ID get-profile
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) me(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := t.service.GetAccount(accountId)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Router /api/Account/Update [put]
// @Security ApiKeyAuth
// @Summary UpdateProfile
// @Tags Account
// @Description Update Profile
// @ID update-profile
// @Accept json
// @Produce json
// @Param Input body models.AccountRequest true "Update Info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) update(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req models.AccountRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.UpdateAccount(accountId, req); err != nil {
		if errors.Is(err, custom_errors.ErrUsernameExists) || errors.Is(err, custom_errors.ErrInvalidParams) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "account successfully updated"})
}
