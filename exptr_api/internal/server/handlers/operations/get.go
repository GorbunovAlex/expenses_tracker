package operations

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/models"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetOperationHandler interface {
	GetOperationsByUserID(userID int) ([]models.Operation, error)
	GetUserIDByToken(token string) (int, error)
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

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		token := r.Header.Get("Bearer")
		userID, err := getAllOperationHandler.GetUserIDByToken(token)
		if err != nil {
			log.Error("failed to get user id by token", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		operations, err := getAllOperationHandler.GetOperationsByUserID(userID)
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
