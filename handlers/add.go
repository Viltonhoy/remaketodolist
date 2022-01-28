package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// структура для анмаршелинга
type JsAct struct {
	St string
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var typ JsAct

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &typ)

	if err != nil {
		fmt.Println("Error", err)
	}

	if typ.St == "" {
		http.Error(w, "Missing Field \"St\"", http.StatusBadRequest)
		return
	}

	h.Counter++
	h.Storage[h.Counter] = typ.St
	js, err := json.Marshal(h.Storage)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Write(js)

}
