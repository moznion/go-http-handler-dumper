package hhdumper

import (
	"fmt"
	"net/http"
)

func ExampleDump() {
	http.Handle("/foo", http.HandlerFunc(foo))
	http.HandleFunc("/bar/buz", func(w http.ResponseWriter, r *http.Request) {
	})

	routes := Dump()

	fmt.Printf("%v\n", routes)
}

func ExampleDumpBy() {
	var mux http.ServeMux
	mux.Handle("/foo", http.HandlerFunc(foo))
	mux.HandleFunc("/bar/buz", func(w http.ResponseWriter, r *http.Request) {
	})

	routes := DumpBy(&mux)

	fmt.Printf("%v\n", routes)
}
