package middleware

import (
	"errors"
	"strings"
)

func parseBearerToken(token string) (string, error) {
	var accessToken string
	accessToken = strings.Replace(token, "Bearer ", "", 1)
	if accessToken == "" {
		return "", errors.New("An invalid access token was provided")
	}

	return accessToken, nil
}
