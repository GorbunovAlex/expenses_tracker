package categories

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type DeleteCategoryHandler interface {
	DeleteCategory(id uuid.UUID) error
}

// Delete godoc
// @Summary      Delete category by id
// @Description  Delete category by id
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id path string true "Category ID"
// @Success      200  {object}  response.Response
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /categories/{id} [delete]
func Delete(log *slog.Logger, deleteCategoryHandler DeleteCategoryHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.categories.delete.DeleteCategory"

		r := c.Request
		w := c.Writer

		param := c.Param("id")
		if param == "" {
			log.Error("empty category id")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("empty category id"))
			return
		}

		id, err := uuid.Parse(param)
		if err != nil {
			log.Error("invalid id format", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("invalid id format"))
			return
		}

		err = deleteCategoryHandler.DeleteCategory(id)
		if err != nil {
			log.Error("failed to delete category", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to delete category"))
			return
		}

		log.Info("category deleted")
		render.JSON(w, r, response.OK())
	}
}
