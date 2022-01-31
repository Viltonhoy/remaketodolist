package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Testadd(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://loacalhost:9090/add", nil)
	w := httptest.NewRecorder()

}
