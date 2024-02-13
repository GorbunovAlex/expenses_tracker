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

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type LoginHandler interface {
	GetUserByEmail(email string) (*models.User, error)
	SetUserSession(userID int, token string) error
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
