package users

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"alex_gorbunov_exptr_api/pkg/hasher"
	"alex_gorbunov_exptr_api/pkg/jwt"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type LoginHandler interface {
	GetUserByEmail(email string) (*models.User, error)
}

func Login(log *slog.Logger, loginHandler LoginHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.users.login.Login"

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

		log.Info("user logged in", slog.Any("user", user))

		render.JSON(w, r, models.LoginResponse{
			Token:    token,
			Response: response.OK(),
		})
	}
}
