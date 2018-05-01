package phone_test

import (
	"testing"
	"github.com/arhitiron/countries/phone"
	"fmt"
)

const (
	usCountryCode      = "US"
	invalidCountryCode = "MOCK_COUNTRY_CODE"
	usPhoneCode        = "1"
)

func TestGetPhoneValid(t *testing.T) {
	phoneCode, err := phone.GetPhone(usCountryCode)
	if err != nil {
		t.Fail()
	}

	if phoneCode != usPhoneCode {
		t.Fail()
	}
}

func TestGetPhoneInValid(t *testing.T) {
	phoneCode, err := phone.GetPhone(invalidCountryCode)
	if err == nil {
		t.Fail()
	}

	if err.Error() != fmt.Sprintf(phone.ErrorFormat, invalidCountryCode) {
		t.Fail()
	}

	if phoneCode == usPhoneCode {
		t.Fail()
	}
}
