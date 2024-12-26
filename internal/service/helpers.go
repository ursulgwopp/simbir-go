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

// ACCOUNT VALIDATIONS

func validateAccountId(s *Service, accountId int) error {
	exists, err := s.repo.CheckAccountIdExists(accountId)

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

func validateUsernameUniqueness(s *Service, username string) error {
	exists, err := s.repo.CheckUsernameExists(username)

	if err != nil {
		return err
	}

	if exists {
		return custom_errors.ErrUsernameIsNotUnique
	}

	return nil
}

func validateUpdatedUsernameUniqueness(s *Service, accountId int, username string) error {
	equal, err := s.repo.CheckUsernameIsEqualToOld(accountId, username)
	if err != nil {
		return err
	}

	if !equal {
		exists, err := s.repo.CheckUsernameExists(username)
		if err != nil {
			return err
		}

		if exists {
			return custom_errors.ErrUsernameIsNotUnique
		}
	}

	return nil
}

func validateAccountDeletion(s *Service, accountId int) error {
	has, err := s.repo.CheckAccountIdHasActiveRents(accountId)

	if err != nil {
		return err
	}

	if has {
		return custom_errors.ErrCanNotDeleteAccount
	}

	return nil
}

// TRANSPORT VALIDATIONS

func validateTransportId(s *Service, transportId int) error {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrTransportIdNotFound
	}

	return nil
}

func validateTransportOwner(s *Service, userId int, transportId int) error {
	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err != nil {
		return err
	}

	if userId != ownerId {
		return custom_errors.ErrAccessDenied
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

func validateTransportDeletion(s *Service, transportId int) error {
	has, err := s.repo.CheckTransportIdHasActiveRents(transportId)
	if err != nil {
		return err
	}

	if has {
		return custom_errors.ErrCanNotDeleteTransport
	}

	return nil
}

func validateTransportIsAvailable(s *Service, transportId int) error {
	is_available, err := s.repo.CheckTransportIsAvailable(transportId)
	if err != nil {
		return err
	}

	if !is_available {
		return custom_errors.ErrTransportNotAvailable
	}

	return nil
}

// RENT VALIDATIONS

func validateRentId(s *Service, rentId int) error {
	exists, err := s.repo.CheckRentIdExists(rentId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrTransportIdNotFound
	}

	return nil
}

func validateRentAccess(s *Service, userId int, transportId int, rentOwnerId int) error {
	transportOwnerId, err := s.repo.CheckOwnerId(transportId)
	if err != nil {
		return err
	}

	if userId != transportOwnerId && userId != rentOwnerId {
		return custom_errors.ErrAccessDenied
	}

	return nil
}

func validateRentOwner(s *Service, userId int, rentId int) error {
	ownerId, err := s.repo.CheckRentOwnerId(rentId)
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

func validateRentIsActive(s *Service, rentId int) error {
	is_active, err := s.repo.CheckRentIsActive(rentId)
	if err != nil {
		return err
	}

	if !is_active {
		return custom_errors.ErrAlreadyStopped
	}

	return nil
}
