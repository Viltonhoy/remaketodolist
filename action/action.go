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
	Action map[int]JsAct
}

func (h *Act) action(w http.ResponseWriter, r *http.Request) {
	var typ JsAct

	str, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(str, &typ)

	if err != nil {
		fmt.Println("Error", err)
	}

	if typ.St == "" {
		http.Error(w, "Missing Field \"St\"", http.StatusBadRequest)
	}

}
