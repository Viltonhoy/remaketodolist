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

func TestHandlerRewrite(t *testing.T) {
	t.Run("green case", func(t *testing.T) {
		arg := bytes.NewBuffer([]byte(`{"Num":2, "Str":"lala"}`))
		req := httptest.NewRequest(http.MethodPost, "http://loacalhost:9090/add", arg)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: map[int]string{1: "da", 2: "ba", 3: "dab"},
			Counter: 3,
		}

		s.Rewrite(w, req)

		bodytest := []byte(`{"1":"da","2":"lala","3":"dab"}`)
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		assert.Equal(t, bodytest, body)
	})

	t.Run("empty body request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://loacalhost:9090/add", nil)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: map[int]string{1: "da", 2: "ba", 3: "dab"},
			Counter: 3,
		}

		s.Rewrite(w, req)

		resp, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "Missing Fields \"Num\", \"Str\"\n", string(resp))
	})
}
