package categories

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
	"github.com/go-playground/validator"
)

type CreateCategoryHandler interface {
	CreateCategory(models.CategoryRequest) error
}

// Create godoc
// @Summary      create new category
// @Description  create new category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        data body models.CategoryRequest true "Update category"
// @Success      200  {object}  models.CategoryResponse
// @Failure      400  {string} 	string "empty request body"
// @Failure      500  {string}  string "server error"
// @Router       /categories/new [post]
func New(log *slog.Logger, createCategoryHandler CreateCategoryHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.categories.create.CreateCategory"

		r := c.Request
		w := c.Writer

		var req models.CategoryRequest

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

		err = createCategoryHandler.CreateCategory(req)

		if err != nil {
			log.Error("failed to create category", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to create category"))
			return
		}

		log.Info("category created", slog.Any("category", req))

		render.JSON(w, r, models.CategoryResponse{
			Response: response.OK(),
		})
	}
}
