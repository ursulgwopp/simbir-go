package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"regexp"

	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
)

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}

func validateUsername(username string) error {
	if len(username) < 3 || len(username) > 30 {
		return custom_errors.ErrInvalidParams
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username); !matched {
		return custom_errors.ErrInvalidParams
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 3 {
		return custom_errors.ErrInvalidParams
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
