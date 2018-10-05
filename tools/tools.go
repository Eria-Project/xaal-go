package tools

import "regexp"

const xaalAddrPattern = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"

func IsValidAddr(val string) bool {
	if val == "" {
		return false
	}
	re := regexp.MustCompile(xaalAddrPattern)
	if re.MatchString(val) {
		return true
	}
	return false
}
