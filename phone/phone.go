package phone

import "fmt"

const ErrorFormat = "Invalid country code %v"

func GetPhone(countryCode string) (string, error) {
	phoneCode := data[countryCode]
	if phoneCode == "" {
		return "", fmt.Errorf(ErrorFormat, countryCode)
	}
	return phoneCode, nil
}

func GetPhoneFormatted(countryCode string) (string, error) {
	phoneCode, err := GetPhone(countryCode)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("+%v", phoneCode), nil
}
