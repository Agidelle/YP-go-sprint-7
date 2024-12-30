package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerStatusCodeAndNotEmptyBody(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(res, req)

	require.Equalf(t, res.Code, http.StatusOK, "expected status code: %d, got %d", http.StatusOK, res.Code)
	assert.NotEmpty(t, res.Body, "body empty")
}

func TestMainHandlerCheckNotActualCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=belgorod", nil)

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(res, req)

	require.Equalf(t, res.Code, http.StatusBadRequest, "excepted status code: %d, got %d", http.StatusBadRequest, res.Code)
	assert.Equalf(t, res.Body.String(), "wrong city value", "expected response: %s, got %s", "wrong city value", res.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(res, req)

	require.Equalf(t, res.Code, http.StatusOK, "expected status code: %d, got %d", http.StatusOK, res.Code)
	list := strings.Split(res.Body.String(), ",")
	assert.Lenf(t, list, totalCount, "expected cafe count: %d, got %d", totalCount, len(list))
}
