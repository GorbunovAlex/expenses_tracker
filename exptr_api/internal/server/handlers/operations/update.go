package operations

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type UpdateOperationHandler interface {
	UpdateOperation(id uuid.UUID, operation *models.OperationRequest) error
}

// UpdateOperation godoc
// @Summary      Update operation by id
// @Description  Update operation by id
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param        id path string true "operation id" data body models.OperationRequest
// @Success      200  {object}  models.UpdateOperationResponse
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /operations/{id} [put]
func Update(log *slog.Logger, updateOperationHandler UpdateOperationHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.operations.updated.UpdateOperation"

		r := c.Request
		w := c.Writer

		param := c.Params.ByName("id")
		if param == "" {
			log.Error("empty id")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("empty id"))
			return
		}

		id, err := uuid.Parse(param)
		if err != nil {
			log.Error("invalid id format", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("invalid id format"))
			return
		}

		var req models.OperationRequest

		err = render.DecodeJSON(r.Body, &req)

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

		err = updateOperationHandler.UpdateOperation(id, &req)
		if err != nil {
			log.Error("failed to update operation", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to update operation"))
			return
		}

		log.Info("operation updated")
		render.JSON(w, r, models.UpdateOperationResponse{
			Response: response.OK(),
		})
	}
}
