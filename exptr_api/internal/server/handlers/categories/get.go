package categories

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
)

type GetCategoriesHandler interface {
	GetCategories(userID int) ([]models.Category, error)
	GetUserIDByToken(token string) (int, error)
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
		const op = "handlers.categories.get.GetCategories"

		r := c.Request
		w := c.Writer

		token := r.Header.Get("Bearer")
		userID, err := getAllCategoriesHandler.GetUserIDByToken(token)
		if err != nil {
			log.Error("failed to get user id by token", sl.Error(err))
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
