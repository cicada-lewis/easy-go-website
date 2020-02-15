package utils

import (
	"github.com/google/uuid"
)

func NewUUID() (string, error) {

	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}