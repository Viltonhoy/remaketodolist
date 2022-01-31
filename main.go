package main

import (
	"log"
	"net/http"
	"remaketodolist/handlers"
)

func main() {

	h := handlers.Handler{
		Storage: make(map[int]string),
	}

	http.HandleFunc("/add", h.Add)
	http.HandleFunc("/del", h.Delete)

	port := ":9090"
	err := http.ListenAndServe(port, nil)
	//принимает двапараметра — порт сединения и функцию-обработчик,
	//которая будет выполнена при запуке сервера.
	if err != nil {
		log.Fatal("ListernAndServe", err)
	}
}
