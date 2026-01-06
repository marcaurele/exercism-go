package phonenumber

import (
	"errors"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	nonDigitRegex := regexp.MustCompile(`[^\d]`)
	nanpRegex := regexp.MustCompile(`^[2-9]\d{2}[2-9]\d{6}$`)

	cleaned := nonDigitRegex.ReplaceAllString(phoneNumber, "")

	if len(cleaned) == 11 {
		if cleaned[0] != '1' {
			return "", errors.New("11 digits must start with 1")
		}
		cleaned = cleaned[1:]
	}

	if len(cleaned) != 10 {
		return "", errors.New("must be 10 or 11 digits")
	}

	if !nanpRegex.MatchString(cleaned) {
		return "", errors.New("invalid NANP format")
	}

	return cleaned, nil
}

func AreaCode(phoneNumber string) (string, error) {
	clean, error := Number(phoneNumber)
	if error != nil {
		return clean, error
	}
	return clean[:3], nil
}

func Format(phoneNumber string) (string, error) {
	clean, error := Number(phoneNumber)
	if error != nil {
		return clean, error
	}

	return "(" + clean[:3] + ") " + clean[3:6] + "-" + clean[6:], nil
}
