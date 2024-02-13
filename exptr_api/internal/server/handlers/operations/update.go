package operations

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UpdateOperationHandler interface {
	UpdateOperation(id int, operation *models.OperationRequest) error
}

func Update(log *slog.Logger, updateOperationHandler UpdateOperationHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.operations.updated.UpdateOperation"

		r := c.Request
		w := c.Writer

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		param := c.Params.ByName("id")
		if param == "" {
			log.Error("empty id")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("empty id"))
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			log.Error("failed to convert id", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
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
