package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

// @Router /api/Transport/{id} [get]
// @Summary GetTransport
// @Tags Transport
// @Description Get Transport
// @ID get-transport
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Success 200 {object} models.TransportResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) getTransport(c *gin.Context) {
	transportId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transport, err := t.service.GetTransport(transportId)
	if err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transport)
}

// @Router /api/Transport [post]
// @Security ApiKeyAuth
// @Summary CreateTransport
// @Tags Transport
// @Description Create Transport
// @ID create-transport
// @Accept json
// @Produce json
// @Param Input body models.TransportRequest true "Transport Info"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) createTransport(c *gin.Context) {
	userId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req models.TransportRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.CreateTransport(userId, req)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
}

// @Router /api/Transport/{id} [put]
// @Security ApiKeyAuth
// @Summary UpdateTransport
// @Tags Transport
// @Description Update Transport
// @ID update-transport
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Param Input body models.TransportRequest true "Transport Info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) updateTransport(c *gin.Context) {
	userId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transportId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req models.TransportRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.UpdateTransport(userId, transportId, req); err != nil {
		if errors.Is(err, custom_errors.ErrAccessDenied) || errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "transport successfully updated"})
}

// @Router /api/Transport/{id} [delete]
// @Security ApiKeyAuth
// @Summary DeleteTransport
// @Tags Transport
// @Description Delete Transport
// @ID delete-transport
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) deleteTransport(c *gin.Context) {
	userId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transportId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.DeleteTransport(userId, transportId); err != nil {
		if errors.Is(err, custom_errors.ErrAccessDenied) || errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "transport successfully deleted"})
}
