package operations

import (
	"alex_gorbunov_exptr_api/internal/lib/logger/handlers/slogdiscard"
	"alex_gorbunov_exptr_api/internal/models"
	"alex_gorbunov_exptr_api/internal/server/handlers/operations/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateOperationHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	categoryID := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	cases := []struct {
		name       string
		input      string
		mockError  error
		setupMock  bool
		statusCode int
		respError  string
	}{
		{
			name: "success",
			input: `{
				"user_id":"11111111-1111-1111-1111-111111111111",
				"category_id":"22222222-2222-2222-2222-222222222222",
				"amount":100,
				"currency":"USD",
				"name":"Test Operation",
				"comment":"test comment",
				"type":"expense",
				"created_at":"2024-01-01T00:00:00Z"
			}`,
			setupMock:  true,
			mockError:  nil,
			statusCode: http.StatusOK,
			respError:  "",
		},
		{
			name:       "empty body",
			input:      "",
			setupMock:  false,
			statusCode: http.StatusBadRequest,
			respError:  "empty request body",
		},
		{
			name:       "invalid json",
			input:      `{invalid}`,
			setupMock:  false,
			statusCode: http.StatusBadRequest,
			respError:  "failed to decode request",
		},
		{
			name: "missing required field",
			input: `{
				"user_id":"11111111-1111-1111-1111-111111111111",
				"amount":100
			}`,
			setupMock:  false,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			createOperationMock := mocks.NewCreateOperationHandler(t)

			if tc.setupMock {
				createOperationMock.On("CreateOperation", mock.MatchedBy(func(op models.OperationRequest) bool {
					return op.UserID == userID && op.CategoryID == categoryID
				})).Return(tc.mockError).Once()
			}

			log := slogdiscard.NewDiscardLogger()
			handler := New(log, createOperationMock)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/operations/new", bytes.NewReader([]byte(tc.input)))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			handler(c)

			require.Equal(t, tc.statusCode, w.Code)

			if tc.respError != "" {
				var resp map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)
				require.Equal(t, tc.respError, resp["error"])
			}
		})
	}
}

func TestCreateOperationHandler_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	categoryID := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	createOperationMock := mocks.NewCreateOperationHandler(t)

	expectedOperation := models.OperationRequest{
		UserID:     userID,
		CategoryID: categoryID,
		Amount:     100,
		Currency:   "USD",
		Name:       "Test Operation",
		Comment:    "test comment",
		Type:       "expense",
		CreatedAt:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	createOperationMock.On("CreateOperation", mock.MatchedBy(func(op models.OperationRequest) bool {
		return op.UserID == expectedOperation.UserID &&
			op.CategoryID == expectedOperation.CategoryID &&
			op.Amount == expectedOperation.Amount &&
			op.Currency == expectedOperation.Currency &&
			op.Name == expectedOperation.Name &&
			op.Comment == expectedOperation.Comment &&
			op.Type == expectedOperation.Type
	})).Return(nil).Once()

	log := slogdiscard.NewDiscardLogger()

	router := gin.New()
	router.POST("/operations/new", New(log, createOperationMock))

	input := `{
		"user_id":"11111111-1111-1111-1111-111111111111",
		"category_id":"22222222-2222-2222-2222-222222222222",
		"amount":100,
		"currency":"USD",
		"name":"Test Operation",
		"comment":"test comment",
		"type":"expense",
		"created_at":"2024-01-01T00:00:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/operations/new", bytes.NewReader([]byte(input)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "OK", resp["status"])
}
