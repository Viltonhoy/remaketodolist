package delete

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type JsAct struct {
	St string
}

type Act struct {
	Action map[int]JsAct
}

type Del struct {
	Id int
}

type Calcul struct {
	Cltr []int
}

func (h *Act) delete(w http.ResponseWriter, r *http.Request) {
	var n Del
	var s Calcul
	s.Cltr = append(s.Cltr, 0)

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &n)

	if err != nil {
		http.Error(w, "Missing Field \"Id\"", http.StatusBadRequest)
		return
	}

	for _, v := range s.Cltr {
		if v == n.Id {
			http.Error(w, "Missing Field \"Id\"", http.StatusBadRequest)
			return
		}
	}

	s.Cltr = append(s.Cltr, n.Id)

	for a := range h.Action {
		if a == n.Id {
			delete(h.Action, n.Id)
		}
	}

	js, err := json.Marshal(h.Action)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
