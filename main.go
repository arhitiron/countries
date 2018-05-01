package main

import (
	"github.com/arhitiron/countries/phone"
	"fmt"
	"regexp"
)

// Simple demo usage
func main() {
	phoneCode, err := phone.GetPhone("US")
	if err != nil {
		panic(err)
	}
	number := 1234567890
	tel := fmt.Sprintf("+%v%v", phoneCode, number)
	var telRegexp = regexp.MustCompile("^\\+" + phoneCode + "([0-9]){3,20}$")
	fmt.Println(telRegexp.MatchString(tel))
}
