package categories

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"alex_gorbunov_exptr_api/internal/domain"
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type CategoryRequest struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Color  string `json:"color"`
	Icon   string `json:"icon"`
}

type UpdateCategoryHandler interface {
	UpdateCategory(category *domain.Category) error
}

// Update godoc
// @Summary      update category
// @Description  update category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id path string true "Category ID" data body models.CategoryRequest true "Update category"
// @Success      200  {object}  response.Response
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /categories/{id} [put]
func Update(log *slog.Logger, updateCategoryHandler UpdateCategoryHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.categories.update.UpdateCategory"

		r := c.Request
		w := c.Writer

		var req CategoryRequest

		id := c.Param("id")
		if id == "" {
			log.Error("empty category id")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("empty category id"))
			return
		}

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

		targetUserUuid, err := uuid.Parse(req.UserID)
		if err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("failed to validate", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error(validateErr.Error()))

			return
		}

		targetCategoryUuid, err := uuid.Parse(id)
		if err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("failed to validate", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error(validateErr.Error()))

			return
		}

		category := domain.Category{
			BaseEntity: domain.BaseEntity{
				ID: targetCategoryUuid,
			},
			UserID: targetUserUuid,
			Name:   req.Name,
			Type:   req.Type,
			Color:  req.Color,
			Icon:   req.Icon,
		}

		err = updateCategoryHandler.UpdateCategory(&category)
		if err != nil {
			log.Error("failed to update category", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to update category"))
			return
		}

		log.Info("category updated", slog.Any("category", req))
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, response.OK())
	}
}
