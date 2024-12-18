package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (t *Transport) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrEmptyAuthHeader.Error())
		return
	}

	valid, err := t.service.CheckTokenIsValid(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !valid {
		models.NewErrorResponse(c, http.StatusUnauthorized, custom_errors.ErrInvalidToken.Error())
		return
	}

	tokenInfo, err := t.service.ParseToken(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("token", header)
	c.Set("account_id", tokenInfo.AccountId)
	c.Set("is_admin", tokenInfo.IsAdmin)
}

func getAccountId(c *gin.Context) (int, error) {
	id, ok := c.Get("account_id")
	if !ok {
		return 0, custom_errors.ErrIdNotFound
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, custom_errors.ErrInvalidIdType
	}

	return idInt, nil
}

func getToken(c *gin.Context) (string, error) {
	token_, ok := c.Get("token")
	if !ok {
		return "", custom_errors.ErrTokenNotFound
	}

	token, ok := token_.(string)
	if !ok {
		return "", custom_errors.ErrInvalidTokenType
	}

	return token, nil
}
