package operations

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
)

type DeleteOperationHandler interface {
	DeleteOperation(id string) error
}

// DeleteOperation godoc
// @Summary      Delete operation by id
// @Description  Delete operation by id
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param        id path string true "operation id"
// @Success      200  {object}  response.Response
// @Failure      400  {string} 	string "empty id"
// @Failure      500  {string}  string "server error"
// @Router       /operations/{id} [delete]
func Delete(log *slog.Logger, deleteOperationHandler DeleteOperationHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.operations.delete.DeleteOperation"

		r := c.Request
		w := c.Writer

		param := c.Params.ByName("id")
		if param == "" {
			log.Error("empty id")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("empty id"))
			return
		}

		err := deleteOperationHandler.DeleteOperation(param)
		if err != nil {
			log.Error("failed to delete operation", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to delete operation"))
			return
		}

		log.Info("operation deleted")
		render.JSON(w, r, response.OK())
	}
}
