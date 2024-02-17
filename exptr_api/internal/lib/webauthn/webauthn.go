package webauthn

import (
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
)

func MustLoad() (*webauthn.WebAuthn, error) {

	const op = "webauthn.MustLoad"

	wconfig := &webauthn.Config{
		RPDisplayName: "ExptrTracker",                                      // Display Name for your site
		RPID:          "http://localhost:3000",                             // Generally the FQDN for your site
		RPOrigins:     []string{"http://localhost:3000/api/v1/users/auth"}, // The origin URLs allowed for WebAuthn requests
	}

	if webAuthn, err := webauthn.New(wconfig); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	} else {
		return webAuthn, nil
	}

}
