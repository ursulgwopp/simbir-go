package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

// @Router /api/Rent/Transport [get]
// @Summary GetAvailableTransport
// @Tags Rent
// @Description Get Available Transport
// @ID get-available-transport
// @Accept json
// @Produce json
// @Param latitude query float64 true "Latitude"
// @Param longitude query float64 true "Longitude"
// @Param radius query float64 true "Radius"
// @Param transportType query string true "TransportType"
// @Success 200 {array} models.AdminTransportResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) getAvailableTransport(c *gin.Context) {
	latitude_ := c.Query("latitude")
	longitude_ := c.Query("longitude")
	radius_ := c.Query("radius")
	transportType := c.Query("transportType")

	latitude, err := strconv.ParseFloat(latitude_, 64)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	longitude, err := strconv.ParseFloat(longitude_, 64)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	radius, err := strconv.ParseFloat(radius_, 64)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transports, err := t.service.GetAvailableTransport(latitude, longitude, radius, transportType)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transports)
}

// @Router /api/Rent/{id} [get]
// @Security ApiKeyAuth
// @Summary GetRent
// @Tags Rent
// @Description Get Rent
// @ID get-rent
// @Accept json
// @Produce json
// @Param id path int true "Rent ID"
// @Success 200 {object} models.RentResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) getRent(c *gin.Context) {
	rentId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rent, err := t.service.GetRent(userId, rentId)
	if err != nil {
		if errors.Is(err, custom_errors.ErrAccessDenied) || errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rent)
}

// @Router /api/Rent/MyHistory [get]
// @Security ApiKeyAuth
// @Summary GetMyHistory
// @Tags Rent
// @Description Get My History
// @ID get-my-history
// @Accept json
// @Produce json
// @Success 200 {array} models.RentResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) getUserHistory(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rents, err := t.service.GetUserHistory(accountId)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rents)
}

// @Router /api/Rent/TransportHistory/{id} [get]
// @Security ApiKeyAuth
// @Summary GetTransportHistory
// @Tags Rent
// @Description Get Transport History
// @ID get-transport-history
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Success 200 {array} models.RentResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) getTransportHistory(c *gin.Context) {
	transportId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getAccountId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rents, err := t.service.GetTransportHistory(userId, transportId)
	if err != nil {
		if errors.Is(err, custom_errors.ErrAccessDenied) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rents)
}

// @Router /api/Rent/New/{id} [post]
// @Security ApiKeyAuth
// @Summary StartRent
// @Tags Rent
// @Description Start Rent
// @ID start-rent
// @Accept json
// @Produce json
// @Param id path int true "Transport ID"
// @Param rentType query string true "Rent Type"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) startRent(c *gin.Context) {
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

	rentType := c.Query("rentType")

	id, err := t.service.StartRent(userId, transportId, rentType)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
}

// @Router /api/Rent/Stop/{id} [post]
// @Security ApiKeyAuth
// @Summary StopRent
// @Tags Rent
// @Description Stop Rent
// @ID stop-rent
// @Accept json
// @Produce json
// @Param id path int true "Rent ID"
// @Param latitude query float64 true "Latitude"
// @Param longitude query float64 true "Longitude"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) stopRent(c *gin.Context) {
	rentId, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	latitude_ := c.Query("latitude")
	longitude_ := c.Query("longitude")

	latitude, err := strconv.ParseFloat(latitude_, 64)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	longitude, err := strconv.ParseFloat(longitude_, 64)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.StopRent(rentId, latitude, longitude); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "rent successfully stopped"})
}
