package categories

import (
	"log/slog"
	"net/http"

	"alex_gorbunov_exptr_api/internal/domain"
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"alex_gorbunov_exptr_api/internal/server/middleware/token"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type GetCategoriesHandler interface {
	GetCategories(userID uuid.UUID) ([]domain.Category, error)
}

// GetAll godoc
// @Summary      get all categories
// @Description  get all categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetCategoriesResponse
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /categories/ [get]
func GetAll(log *slog.Logger, getAllCategoriesHandler GetCategoriesHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const fn = "handlers.categories.get.GetCategories"
		log = log.With(slog.String("fn", fn))

		r := c.Request
		w := c.Writer

		userIDStr, ok := token.GetUserIDFromContext(c)
		if !ok {
			log.Error("failed to get user id from context")
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, response.Error("unauthorized"))
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			log.Error("failed to parse user id", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		categories, err := getAllCategoriesHandler.GetCategories(userID)
		if err != nil {
			log.Error("failed to get all categories", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get all categories"))
			return
		}

		log.Info("all categories received")
		render.JSON(w, r, models.GetCategoriesResponse{
			Response:   response.OK(),
			Categories: categories,
		})
	}
}
