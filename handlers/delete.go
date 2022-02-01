package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type del struct {
	Id int
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var n del

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &n)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

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
