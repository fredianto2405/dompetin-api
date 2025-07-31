package password

import (
	"errors"
	"regexp"
)

func Validate(password string) error {
	if len(password) < 8 {
		return errors.New(MsgPasswordMin)
	}

	lowercase := regexp.MustCompile(`[a-z]`)
	uppercase := regexp.MustCompile(`[A-Z]`)
	number := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[@#$_&\-+]`)

	switch {
	case !lowercase.MatchString(password):
		return errors.New(MsgPasswordLower)
	case !uppercase.MatchString(password):
		return errors.New(MsgPasswordUpper)
	case !number.MatchString(password):
		return errors.New(MsgPasswordNumeric)
	case !special.MatchString(password):
		return errors.New(MsgPasswordSpecialChar)
	}

	return nil
}
