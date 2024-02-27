package users

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"alex_gorbunov_exptr_api/pkg/hasher"
	"alex_gorbunov_exptr_api/pkg/jwt"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	authn "alex_gorbunov_exptr_api/internal/lib/wauthn"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/go-webauthn/webauthn/webauthn"
)

type LoginHandler interface {
	GetUserByEmail(email string) (*models.User, error)
	SetUserSession(userID int, token string) error
	SetAuthnUserSession(userID int, session *webauthn.SessionData) error
	GetAuthnUserSession(userID int) (*webauthn.SessionData, error)
}

// Login godoc
// @Summary      Login
// @Description  Login
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        data body  models.LoginRequest  true  "login request"
// @Success      200  {object}  models.LoginResponse
// @Failure      400  {string} 	string "empty request body"
// @Failure      404  {string}  string "wrong email or password"
// @Failure      500  {string}  string "server error"
// @Router       /users/login [post]
func Login(log *slog.Logger, loginHandler LoginHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.users.login.Login"

		r := c.Request
		w := c.Writer

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		var req models.LoginRequest

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
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error(validateErr.Error()))

			return
		}

		user, err := loginHandler.GetUserByEmail(req.Email)
		if err != nil {
			log.Error("failed to get user by email", sl.Error(err))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, response.Error("wrong email or password"))
			return
		}

		if !hasher.CheckPasswordHash(req.Password, user.Password) {
			log.Error("passwords do not match")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("passwords do not match"))
			return
		}

		token, err := jwt.GetSignedToken()
		if err != nil {
			log.Error("failed to get signed token", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		err = loginHandler.SetUserSession(user.ID, token)
		if err != nil {
			log.Error("failed to set user session", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		log.Info("user logged in", slog.Any("user", user))

		render.JSON(w, r, models.LoginResponse{
			Token:    token,
			Response: response.OK(),
		})
	}
}

func LoginWebAuthnBegin(log *slog.Logger, loginHandler LoginHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.users.login.LoginWebAuthnBegin"

		r := c.Request
		w := c.Writer

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		var req models.WebAuthnUserRequest

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
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error(validateErr.Error()))

			return
		}

		user, err := loginHandler.GetUserByEmail(req.Email)
		if err != nil {
			log.Error("failed to get user by email", sl.Error(err))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, response.Error("wrong email or password"))
			return
		}

		options, session, err := authn.WebAuthn.BeginLogin(user)
		if err != nil {
			log.Error("failed to begin login", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		err = loginHandler.SetAuthnUserSession(user.ID, session)
		if err != nil {
			log.Error("failed to set user session", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		render.JSON(w, r, models.WebAuthnLoginResponse{
			Options:  options,
			Response: response.OK(),
		})

	}
}

func LoginWebAuthnFinish(log *slog.Logger, loginHandler LoginHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.users.login.LoginWebAuthnFinish"

		r := c.Request
		w := c.Writer

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		var req models.WebAuthnUserRequest

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
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error(validateErr.Error()))

			return
		}

		user, err := loginHandler.GetUserByEmail(req.Email)
		if err != nil {
			log.Error("failed to get user by email", sl.Error(err))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, response.Error("wrong email or password"))
			return
		}

		session, err := loginHandler.GetAuthnUserSession(user.ID)
		if err != nil {
			log.Error("failed to get user session", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		credential, err := authn.WebAuthn.FinishLogin(user, *session, r)
		if err != nil {
			log.Error("failed to finish login", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		credentialJson, err := json.Marshal(credential)
		if err != nil {
			log.Error("failed to marshal credential", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		credentialData := string(credentialJson)

		err = loginHandler.SetUserSession(user.ID, credentialData)
		if err != nil {
			log.Error("failed to set user session", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		render.JSON(w, r,
			models.LoginResponse{
				Token:    credentialData,
				Response: response.OK(),
			})
	}
}
