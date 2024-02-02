package operations

import (
	"alex_gorbunov_exptr_api/internal/http-server/handlers/operations/mocks"
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/internal/lib/logger/handlers/slogdiscard"
	"alex_gorbunov_exptr_api/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateOperationHandler(t *testing.T) {
	cases := []struct {
		name      string
		operation models.CreateOperationRequest
		respError string
		mockError error
	}{
		{
			name: "success",
			operation: models.CreateOperationRequest{
				UserID:     "user_id",
				CategoryID: "category_id",
				Amount:     100,
				Currency:   "USD",
				Name:       "name",
				Comment:    "comment",
				Type:       "type",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			createOperationMock := mocks.NewCreateOperationHandler(t)
			if tc.respError == "" || tc.mockError != nil {
				createOperationMock.On("CreateOperation", tc.operation, mock.AnythingOfType("context.Context")).Return(&models.CreateOperationResponse{}, tc.mockError).Once()
			}

			handler := New(slogdiscard.NewDiscardLogger(), createOperationMock)

			input := `{"user_id":"user_id","category_id":"category_id","amount":100,"currency":"USD","name":"name","comment":"comment","type":"type"}`

			req, err := http.NewRequest(http.MethodPost, "/operations", bytes.NewReader([]byte(input)))
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			require.Equal(t, rr.Code, http.StatusOK)
			body := rr.Body.String()
			var resp response.Response
			require.NoError(t, json.Unmarshal([]byte(body), &resp))
			require.Equal(t, tc.respError, resp.Error)
		})
	}
}
