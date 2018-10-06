package tools

import (
	"regexp"

	"github.com/satori/go.uuid"
)

const xaalAddrPattern = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"

// IsValidAddr : Check is the xAAL address is valid
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

// GetRandomUUID : Generates a new xAAL UUID
func GetRandomUUID() string {
	u1 := uuid.Must(uuid.NewV1()) // panic on error
	return u1.String()
}
