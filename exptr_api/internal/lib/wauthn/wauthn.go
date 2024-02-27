package wauthn

import (
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	WebAuthn *webauthn.WebAuthn
)

func MustLoad() error {

	const op = "webauthn.MustLoad"

	wconfig := &webauthn.Config{
		RPDisplayName: "ExptrTracker",                                      // Display Name for your site
		RPID:          "http://localhost:3000",                             // Generally the FQDN for your site
		RPOrigins:     []string{"http://localhost:3000/api/v1/users/auth"}, // The origin URLs allowed for WebAuthn requests
	}

	if webAuthn, err := webauthn.New(wconfig); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	} else {
		WebAuthn = webAuthn
		return nil
	}

}
