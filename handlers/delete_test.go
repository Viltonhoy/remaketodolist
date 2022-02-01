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

func TestHandlerDelete(t *testing.T) {
	t.Run("green case", func(t *testing.T) {

		arg := bytes.NewBuffer([]byte(`{"Id":2}`))
		req := httptest.NewRequest(http.MethodGet, "http://loacalhost:9090/del", arg)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: map[int]string{1: "da", 2: "ba", 3: "dab"},
			Counter: 3,
		}

		s.Delete(w, req)

		bodytest := []byte(`{"1":"da","3":"dab"}`)
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		assert.Equal(t, bodytest, body)
	})

	t.Run("empty request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://loacalhost:9090/del", nil)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: map[int]string{1: "da", 2: "ba", 3: "dab"},
			Counter: 3,
		}

		s.Delete(w, req)

		resp, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "Empty request body\n", string(resp))
	})

	t.Run("epty Id", func(t *testing.T) {
		arg := bytes.NewBuffer([]byte(`{"Id":4}`))
		req := httptest.NewRequest(http.MethodGet, "http://loacalhost:9090/del", arg)
		w := httptest.NewRecorder()

		s := Handler{
			Storage: map[int]string{1: "da", 2: "ba", 3: "dab"},
			Counter: 3,
		}

		s.Delete(w, req)

		resp, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "Wrong Id\n", string(resp))
	})
}
