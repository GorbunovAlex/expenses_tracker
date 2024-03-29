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
	"github.com/go-playground/validator/v10"
)

//go:generate mockery --name=CreateOperationHandler
type CreateOperationHandler interface {
	CreateOperation(models.OperationRequest) error
}

// New godoc
// @Summary      Create new operation
// @Description  Create new operation
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param        data body models.OperationRequest  true  "Create operation"
// @Success      200  {object}  models.CreateOperationResponse
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /operations/new [post]
func New(log *slog.Logger, createOperationHandler CreateOperationHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.operations.create.CreateOperation"

		r := c.Request
		w := c.Writer

		var req models.OperationRequest

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

		err = createOperationHandler.CreateOperation(req)

		if err != nil {
			log.Error("failed to create operation", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to create operation"))
			return
		}

		log.Info("operation created", slog.Any("operation", req))

		render.JSON(w, r, models.CreateOperationResponse{
			Response: response.OK(),
		})
	}
}
