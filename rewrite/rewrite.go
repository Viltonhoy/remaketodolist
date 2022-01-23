package rewrite

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Act struct {
	Action  map[int]JsAct
	counter int
}

type JsAct struct {
	St string
}

type Rew struct {
	Num int
	Str string
}

func (h *Act) rewrite(w http.ResponseWriter, r *http.Request) {
	var mp Rew

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &mp)

	if err != nil {
		http.Error(w, "Missing Fields \"Num\", \"Str\"", http.StatusBadRequest)
		return
	}

	for a := range h.Action {
		if a == mp.Num {
			h.Action[a] = JsAct{mp.Str}
		}
	}

	js, err := json.Marshal(h.Action)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
