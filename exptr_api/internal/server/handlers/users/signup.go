package users

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	authn "alex_gorbunov_exptr_api/internal/lib/wauthn"
	"alex_gorbunov_exptr_api/internal/models"
	"alex_gorbunov_exptr_api/pkg/hasher"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/go-webauthn/webauthn/webauthn"
)

type SignupHandler interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	SetAuthnUserSession(userID int, session *webauthn.SessionData) error
	GetAuthnUserSession(userID int) (*webauthn.SessionData, error)
	SetAuthnUserCredentials(userID int, credentials *webauthn.Credential) error
}

// Signup godoc
// @Summary      Signup
// @Description  Signup
// @Tags         users
// @Accept       json
// @Produce      json
// @Param				 data body  models.SignUpRequest  true  "signup request"
// @Success      200  {object}  response.Response
// @Failure      400  {string}  string "empty request body"
// @Failure      404  {string}  string "user already exists"
// @Failure      500  {string}  string "server error"
// @Router       /users/signup [post]
func Signup(log *slog.Logger, signupHandler SignupHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.users.signup.Signup"

		r := c.Request
		w := c.Writer

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		var req models.SignUpRequest

		err := render.DecodeJSON(r.Body, &req)

		if errors.Is(err, io.EOF) {
			log.Error("empty request body")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("empty request body"))
			return
		}

		if err != nil {
			log.Error("failed to decode request", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}

		log.Info("request decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("validation failed", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error(validateErr.Error()))

			return
		}

		_, err = signupHandler.GetUserByEmail(req.Email)
		if err == nil {
			log.Error("user with this email already exists")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("user already exists"))
			return
		}

		passwordHash, err := hasher.HashPassword(req.Password)
		if err != nil {
			log.Error("failed to hash password", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to hash password"))
			return
		}

		user := &models.User{
			Email:     req.Email,
			Password:  passwordHash,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = signupHandler.CreateUser(user)

		if err != nil {
			log.Error("failed to create user", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to create user"))
			return
		}

		log.Info("user created", slog.Any("user", user))

		render.JSON(w, r, response.OK())
	}
}

func SignupWebAuthnBegin(log *slog.Logger, signupHandler SignupHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.users.signup.SignupWebAuthnBegin"

		r := c.Request
		w := c.Writer

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		var req models.WebAuthnUserRequest

		user := &models.User{
			Email:     req.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		options, session, err := authn.WebAuthn.BeginRegistration(user)

		if err != nil {
			log.Error("failed to begin registration", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to begin registration"))
			return
		}

		err = signupHandler.SetAuthnUserSession(user.ID, session)

		if err != nil {
			log.Error("failed to set user session", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to set user session"))
			return
		}

		log.Info("user created", slog.Any("user", user))

		render.JSON(w, r, models.WebAuthnSignupResponse{
			Options:  options,
			Response: response.OK(),
		})
	}
}

func SignupWebAuthnFinish(log *slog.Logger, signupHandler SignupHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.users.signup.SignupWebAuthnFinish"

		r := c.Request
		w := c.Writer

		var req models.WebAuthnUserRequest

		user, err := signupHandler.GetUserByEmail(req.Email)
		if err != nil {
			log.Error("failed to get user by email", sl.Error(err))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, response.Error("wrong email or password"))
			return
		}

		session, err := signupHandler.GetAuthnUserSession(user.ID)
		if err != nil {
			log.Error("failed to get user session", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get user session"))
			return
		}

		credential, err := authn.WebAuthn.FinishRegistration(user, *session, r)
		if err != nil {
			log.Error("failed to finish registration", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to finish registration"))
			return
		}

		err = signupHandler.SetAuthnUserCredentials(user.ID, credential)
		if err != nil {
			log.Error("failed to set user credentials", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to set user credentials"))
			return
		}

		err = signupHandler.CreateUser(user)
		if err != nil {
			log.Error("failed to create user", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to create user"))
			return
		}

		render.JSON(w, r, response.OK())
	}
}
