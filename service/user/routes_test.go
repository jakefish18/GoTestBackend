package user

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jakefish18/GolangHttpServer/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "",
			Password:  "123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handlerRegiser)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got: %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}
