package govimeo

import (
	"errors"
)

var apiToken string

func SetToken(s string) error {
	if s == "" {
		return errors.New("token can not be empty")
	}

	apiToken = s
	return nil
}
