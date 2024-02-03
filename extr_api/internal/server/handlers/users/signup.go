package users

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"alex_gorbunov_exptr_api/pkg/hasher"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type SignupHandler interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

func Signup(log *slog.Logger, signupHandler SignupHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.users.signup.Signup"

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
