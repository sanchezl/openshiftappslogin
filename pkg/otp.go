package pkg

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func RedHatInternalPassword(secret, prefix string) (string, error) {
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return prefix + code, nil
}
