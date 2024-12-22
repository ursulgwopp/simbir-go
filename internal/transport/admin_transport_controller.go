package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

// @Router /api/Admin/Transport [get]
// @Security ApiKeyAuth
// @Summary ListTransports
// @Tags Admin Transport
// @Description List Transports
// @ID list-transports
// @Accept json
// @Produce json
// @Param from query int true "From"
// @Param count query int true "Count"
// @Param transportType query string true "TransportType"
// @Success 200 {array} models.AdminTransportResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminListTransports(c *gin.Context) {
	from_ := c.Query("from")
	count_ := c.Query("count")
	transportType := c.Query("transportType")

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

	transport, err := t.service.AdminListTransports(from, count, transportType)
	if err != nil {
		if errors.Is(err, custom_errors.ErrInvalidParams) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transport)
}

// @Router /api/Admin/Transport/{id} [get]
// @Security ApiKeyAuth
// @Summary GetTransport
// @Tags Admin Transport
// @Description Get Transport
// @ID admin-get-transport
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Success 200 {object} models.AdminTransportResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminGetTransport(c *gin.Context) {
	transportId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transport, err := t.service.AdminGetTransport(transportId)
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

// @Router /api/Admin/Transport [post]
// @Security ApiKeyAuth
// @Summary CreateTransport
// @Tags Admin Transport
// @Description Create Transport
// @ID create-transport
// @Accept json
// @Produce json
// @Param Input body models.AdminTransportRequest true "Transport Info"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminCreateTransport(c *gin.Context) {
	var req models.AdminTransportRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.AdminCreateTransport(req)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
}

// @Router /api/Admin/Transport/{id} [put]
// @Security ApiKeyAuth
// @Summary UpdateTransport
// @Tags Admin Transport
// @Description Update Transport
// @ID update-transport
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Param Input body models.AdminTransportRequest true "Update Info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminUpdateTransport(c *gin.Context) {
	transportId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req models.AdminTransportRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.AdminUpdateTransport(transportId, req); err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "transport successfully updated"})
}

// @Router /api/Admin/Transport/{id} [delete]
// @Security ApiKeyAuth
// @Summary DeleteTransport
// @Tags Admin Transport
// @Description Delete Transport
// @ID admin-delete-transport
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminDeleteTransport(c *gin.Context) {
	transportId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.AdminDeleteTransport(transportId); err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "transport successfully deleted"})
}
