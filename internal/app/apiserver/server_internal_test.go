package apiserver

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSever_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
