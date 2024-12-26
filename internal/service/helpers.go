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

func validatePagination(from int, count int) error {
	if from < 0 || count < 0 {
		return custom_errors.ErrInvalidPaginationParams
	}

	return nil
}

func validateAccountId(exists bool, err error) error {
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrAccountIdNotFound
	}

	return nil
}

func validateAccountRequest(req models.AccountRequest) error {
	if len(req.Username) < 3 || len(req.Username) > 30 {
		return custom_errors.ErrInvalidUsernameLength
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, req.Username); !matched {
		return custom_errors.ErrInvalidUsernameCharacters
	}

	if len(req.Password) < 3 || len(req.Password) > 30 {
		return custom_errors.ErrInvalidPasswordLength
	}

	return nil
}

func validateAdminAccountRequest(req models.AdminAccountRequest) error {
	if len(req.Username) < 3 || len(req.Username) > 30 {
		return custom_errors.ErrInvalidUsernameLength
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, req.Username); !matched {
		return custom_errors.ErrInvalidUsernameCharacters
	}

	if len(req.Password) < 3 || len(req.Password) > 30 {
		return custom_errors.ErrInvalidPasswordLength
	}

	if req.Balance < 0 {
		return custom_errors.ErrInvalidBalanceValue
	}

	return nil
}

func validateUsernameUniqueness(exists bool, err error) error {
	if err != nil {
		return err
	}

	if exists {
		return custom_errors.ErrUsernameIsNotUnique
	}

	return nil
}

func validateUpdatedUsernameUniqueness(equal bool, err1 error, exists bool, err2 error) error {
	if err1 != nil {
		return err1
	}

	if !equal {
		if err2 != nil {
			return err2
		}

		if exists {
			return custom_errors.ErrUsernameIsNotUnique
		}
	}

	return nil
}

func validateAccountDeletion(has bool, err error) error {
	if err != nil {
		return err
	}

	if has {
		return custom_errors.ErrCanNotDeleteAccount
	}

	return nil
}

func validateTransportId(exists bool, err error) error {
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrTransportIdNotFound
	}

	return nil
}

func validateTransportOwner(userId int, ownerId int, err error) error {
	if err != nil {
		return err
	}
	if userId != ownerId {
		return custom_errors.ErrAccessDenied
	}

	return nil
}

func validateTransportOwner1(userId int, ownerId int, err error) error {
	if err != nil {
		return err
	}
	if userId == ownerId {
		return custom_errors.ErrCanNotRent
	}

	return nil
}

func validateTransportType(transportType string) error {
	if transportType != "Car" && transportType != "Bike" && transportType != "Scooter" && transportType != "All" {
		return custom_errors.ErrInvalidTransportType
	}

	return nil
}

func validateTransportRequest(req models.TransportRequest) error {
	if len(req.Model) > 255 || len(req.Color) > 255 || len(req.Description) > 255 || len(req.Identifier) > 255 {
		return custom_errors.ErrInvalidTransportProperties
	}

	if req.Latitude < 0 || req.Longitude < 0 || req.MinutePrice < 0 || req.DayPrice < 0 {
		return custom_errors.ErrInvalidTransportProperties
	}

	return nil
}

func validateAdminTransportRequest(req models.AdminTransportRequest) error {
	if len(req.Model) > 255 || len(req.Color) > 255 || len(req.Description) > 255 || len(req.Identifier) > 255 {
		return custom_errors.ErrInvalidTransportProperties
	}

	if req.Latitude < 0 || req.Longitude < 0 || req.MinutePrice < 0 || req.DayPrice < 0 {
		return custom_errors.ErrInvalidTransportProperties
	}

	return nil
}

func validateTransportDeletion(has bool, err error) error {
	if err != nil {
		return err
	}

	if has {
		return custom_errors.ErrCanNotDeleteTransport
	}

	return nil
}

func validateTransportIsAvailable(is_available bool, err error) error {
	if err != nil {
		return err
	}
	if !is_available {
		return custom_errors.ErrTransportNotAvailable
	}

	return nil
}

func validateRentId(exists bool, err error) error {
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrTransportIdNotFound
	}

	return nil
}

func validateRentAccess(userId int, transportOwnerId int, rentOwnerId int, err error) error {
	if err != nil {
		return err
	}

	if userId != transportOwnerId && userId != rentOwnerId {
		return custom_errors.ErrAccessDenied
	}

	return nil
}

func validateRentOwner(userId int, ownerId int, err error) error {
	if err != nil {
		return err
	}

	if userId != ownerId {
		return custom_errors.ErrAccessDenied
	}

	return nil
}

func validateRentType(rentType string) error {
	if rentType != "Minutes" && rentType != "Days" {
		return custom_errors.ErrInvalidParams
	}

	return nil
}

func validateRentIsActive(is_active bool, err error) error {
	if err != nil {
		return err
	}

	if !is_active {
		return custom_errors.ErrAlreadyStopped
	}

	return nil
}
