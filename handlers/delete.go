package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Del struct {
	Id int
}

type Calcul struct {
	Cltr []int
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
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

	_, b := h.Storage[n.Id]
	if b {
		delete(h.Storage, n.Id)
	} else {
		http.Error(w, "Wrong Id", http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(h.Storage)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
