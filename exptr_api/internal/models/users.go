package models

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	authn "github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	ID        int       `json:"id"`
	AuthnID   []byte    `json:"authn_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) WebAuthnID() []byte {
	return user.AuthnID
}

func (user *User) WebAuthnName() string {
	return "newUser"
}

func (user *User) WebAuthnDisplayName() string {
	return "New User"
}

func (user *User) WebAuthnIcon() string {
	return "https://pics.com/avatar.png"
}

func (user *User) WebAuthnCredentials() []authn.Credential {
	return []authn.Credential{}
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type WebAuthnUserRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type WebAuthnSignupResponse struct {
	Options *protocol.CredentialCreation `json:"options"`
	response.Response
}

type WebAuthnLoginResponse struct {
	Options *protocol.CredentialAssertion `json:"options"`
	response.Response
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	response.Response
}

type UserSession struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}
