package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerAdd(t *testing.T) {
	t.Run("green case", func(t *testing.T) {
		arg := bytes.NewBuffer([]byte(`{"St":"fgfg"}`))
		req := httptest.NewRequest(http.MethodGet, "http://loacalhost:9090/add", arg)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: make(map[int]string),
		}

		s.Add(w, req)
		resptest := `{"1":"fgfg"}`
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		assert.Equal(t, string(resptest), string(body))
	})

	t.Run("empty request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://loacalhost:9090/add", nil)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: make(map[int]string),
		}

		s.Add(w, req)

		resp := w.Result()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), resp.Status)

	})
}
