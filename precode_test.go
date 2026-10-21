
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=moscow", nil)
	rec := httptest.NewRecorder()

	mainHandle(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=paris", nil)
	rec := httptest.NewRecorder()

	mainHandle(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "wrong city value", rec.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=10&city=moscow", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	cafes := strings.Split(rec.Body.String(), ",")
	assert.Len(t, cafes, totalCount)
}
