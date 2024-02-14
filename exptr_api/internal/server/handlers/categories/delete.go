package categories

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
)

type DeleteCategoryHandler interface {
	DeleteCategory(id int) error
}

// Delete godoc
// @Summary      Delete category by id
// @Description  Delete category by id
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param        id path int true "Category ID"
// @Success      200  {object}  response.Response
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /categories/{id} [delete]
func Delete(log *slog.Logger, deleteCategoryHandler DeleteCategoryHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.categories.delete.DeleteCategory"

		r := c.Request
		w := c.Writer

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error("failed to convert id to int", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("server error"))
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
