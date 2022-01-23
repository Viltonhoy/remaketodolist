package main

import (
	"log"
	"net/http"
	"remaketodolist/action"
	//"remaketodolist/delete"
)

func main() {

	h := action.Act{
		ActIon:  make(map[int]action.JsAct),
		Counter: 0,
	}

	http.HandleFunc("/add", h.Action)
	//http.HandleFunc("/del", h.delete.delete)

	port := ":9090"
	err := http.ListenAndServe(port, nil)
	//принимает двапараметра — порт сединения и функцию-обработчик,
	//которая будет выполнена при запуке сервера.
	if err != nil {
		log.Fatal("ListernAndServe", err)
	}
}
