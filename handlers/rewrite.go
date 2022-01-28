package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Rew struct {
	Num int
	Str string
}

func (h *Handler) Rewrite(w http.ResponseWriter, r *http.Request) {
	var mp Rew

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &mp)

	if err != nil {
		http.Error(w, "Missing Fields \"Num\", \"Str\"", http.StatusBadRequest)
		return
	}

	for a := range h.Storage {
		if a == mp.Num {
			h.Storage[a] = mp.Str
		}
	}

	js, err := json.Marshal(h.Storage)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
