package oauth2

import (
	"github.com/outrightmental/go-omniauth/common"
	"github.com/outrightmental/go-objx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationHeader(t *testing.T) {

	creds := &common.Credentials{Map: objx.MSI()}
	accessTokenVal := "ACCESSTOKEN"
	creds.Set(OAuth2KeyAccessToken, accessTokenVal)
	k, v := AuthorizationHeader(creds)

	assert.Equal(t, k, "Authorization")
	assert.Equal(t, "Bearer "+accessTokenVal, v)

}
