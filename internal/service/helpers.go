package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"regexp"

	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}

func validateAccountRequest(req models.AccountRequest) error {
	if len(req.Username) < 3 || len(req.Username) > 30 {
		return custom_errors.ErrInvalidParams
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, req.Username); !matched {
		return custom_errors.ErrInvalidParams
	}

	if len(req.Password) < 3 {
		return custom_errors.ErrInvalidParams
	}

	return nil
}

func validateAdminAccountRequest(req models.AdminAccountRequest) error {
	if len(req.Username) < 3 || len(req.Username) > 30 {
		return custom_errors.ErrInvalidParams
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, req.Username); !matched {
		return custom_errors.ErrInvalidParams
	}

	if len(req.Password) < 3 {
		return custom_errors.ErrInvalidParams
	}

	if req.Balance < 0 {
		return custom_errors.ErrInvalidParams
	}

	return nil
}

func validateUsernameUniqueness(exists bool, err error) error {
	if err != nil {
		return err
	}

	if exists {
		return custom_errors.ErrUsernameExists
	}

	return nil
}

func validateAccountId(exists bool, err error) error {
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	return nil
}

func validateUpdatedUsername(equal bool, err1 error, exists bool, err2 error) error {
	if err1 != nil {
		return err1
	}
	if !equal {
		if err2 != nil {
			return err2
		}

		if exists {
			return custom_errors.ErrUsernameExists
		}
	}

	return nil
}

func validatePagination(from int, count int) error {
	if from < 0 || count < 0 {
		return custom_errors.ErrInvalidParams
	}

	return nil
}

func validateAccountDeletion(has bool, err error) error {
	if err != nil {
		return err
	}

	if has {
		return custom_errors.ErrCanNotDelete
	}

	return nil
}

func validateTransportType(transportType string) error {
	if transportType != "Car" && transportType != "Bike" && transportType != "Scooter" && transportType != "All" {
		return custom_errors.ErrInvalidParams
	}

	return nil
}

func validateTransportProperties(model string, color string, description string, identifier string) error {
	if len(model) > 255 || len(color) > 255 || len(description) > 255 || len(identifier) > 255 {
		return custom_errors.ErrInvalidParams
	}

	return nil
}
