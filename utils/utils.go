package utils

import (
	"encoding/hex"
	"regexp"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/scrypt"
)

const xaalAddrPattern = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"
const xaalDevTypePattern = "^[a-zA-Z][a-zA-Z0-9_-]*\\.[a-zA-Z][a-zA-Z0-9_-]*$"

// IsValidAddr : Check is the xAAL address is valid
func IsValidAddr(val string) bool {
	if val == "" {
		return false
	}
	re := regexp.MustCompile(xaalAddrPattern)
	return re.MatchString(val)
}

// IsValidDevType : Check is the xAAL devType is valid
func IsValidDevType(val string) bool {
	if val == "" {
		return false
	}
	re := regexp.MustCompile(xaalDevTypePattern)
	return re.MatchString(val)
}

// GetRandomUUID : Generates a new xAAL UUID
func GetRandomUUID() string {
	u1 := uuid.Must(uuid.NewV4()) // panic on error
	return u1.String()
}

// Pass2key : Generates key from passphrase using scrypt
func Pass2key(passphrase string) string {
	salt := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // buffer of zeros (crypto_pwhash_scryptsalsa208sha256_SALTBYTES)
	key, _ := scrypt.Key([]byte(passphrase), []byte(salt), 16384, 8, 1, 32)
	return hex.EncodeToString(key)
}

// SliceContains : find a string in a slice
func SliceContains(a *[]string, x string) (int, bool) {
	for i, n := range *a {
		if x == n {
			return i, true
		}
	}
	return -1, false
}

// SliceDelete : delete a slice index
func SliceDelete(a []string, i string) []string {
	if i, in := SliceContains(&a, i); in {
		copy(a[i:], a[i+1:])
		a[len(a)-1] = ""
		a = a[:len(a)-1]
	}
	return a
}
