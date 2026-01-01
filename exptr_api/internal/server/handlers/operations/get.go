package operations

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

type GetOperationHandler interface {
	GetOperationsByUserID(userID uuid.UUID) ([]domain.Operation, error)
}

// GetAll godoc
// @Summary      Get all current user operations
// @Description  Get all current user operations
// @Tags         operations
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetOperationsByUserIDResponse
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /operations [get]
func GetAll(log *slog.Logger, getAllOperationHandler GetOperationHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.operations.get.GetOperationsByUserID"

		r := c.Request
		w := c.Writer

		userIDStr, ok := token.GetUserIDFromContext(c)
		if !ok {
			log.Error("failed to get user id from context")
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, response.Error("unauthorized"))
			return
		}

		targetUserID, err := uuid.Parse(userIDStr)
		if err != nil {
			log.Error("failed to parse user id", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		operations, err := getAllOperationHandler.GetOperationsByUserID(targetUserID)
		if err != nil {
			log.Error("failed to get all operations", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get all operations"))
			return
		}

		log.Info("all operations received")
		render.JSON(w, r, models.GetOperationsByUserIDResponse{
			Response:   response.OK(),
			Operations: operations,
		})
	}
}
