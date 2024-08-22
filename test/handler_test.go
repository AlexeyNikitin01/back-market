package test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func Test_IfGetTokenWithBadGrantType_ShouldReturnBadRequest(t *testing.T) {
	e := httpexpect.Default(t, testServerURL)
	e.POST("/v1/token").WithQuery("grant_type", "authorization_code").
		Expect().Status(http.StatusBadRequest)
}

func Test_IfValidGetToken_ShouldReturnToken(t *testing.T) {
	e := httpexpect.Default(t, testServerURL)
	e.POST("/v1/token").
		WithBytes([]byte("grant_type=password&scope=read+write&username=tst&password=tst")).
		Expect().Status(http.StatusOK).JSON().Object().Keys().ContainsAll("access_token", "expires_in")
}
