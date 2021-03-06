package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
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

		resp, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "Empty request body\n", string(resp))

	})

	t.Run("empty St value", func(t *testing.T) {
		arg := bytes.NewBuffer([]byte(`{"St":""}`))
		req := httptest.NewRequest(http.MethodGet, "http://loacalhost:9090/add", arg)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: make(map[int]string),
		}

		s.Add(w, req)

		resp, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "Missing Field \"St\"\n", string(resp))
	})
}
