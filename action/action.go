package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsAct struct {
	St string
}

type Act struct {
	ActIon  map[int]JsAct
	Counter int
}

func (h *Act) Action(w http.ResponseWriter, r *http.Request) {
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
	h.ActIon[h.Counter] = typ
	js, _ := json.Marshal(h.Action)
	w.Write(js)

}
