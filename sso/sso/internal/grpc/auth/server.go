package auth

import sso "alex_gorbunov/s_pet/protos"

type serverAPI struct {
	sso.UnimplementedAuthServer
}
