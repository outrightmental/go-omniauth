package gomniauth

import (
	"github.com/outrightmental/go-omniauth/common"
)

// SetSecurityKey sets the global security key to be used for signing the state variable
// in the auth request. This allows gomniauth to detect if the data in the
// state variable has been changed.
func SetSecurityKey(key string) {
	common.SetSecurityKey(key)
}

// GetSecurityKey gets the security key.
func GetSecurityKey() string {
	return common.GetSecurityKey()
}
