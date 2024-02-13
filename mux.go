package main

import "net/http"

func setupHandlers(mux *http.ServeMux, requestChannel chan query) {
	mux.HandleFunc("GET /{table}/{key}", func(w http.ResponseWriter, r *http.Request) {
		table := r.PathValue("table")
		key := r.PathValue("key")

		resp, err := getVal(table, key, requestChannel)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Key and/or table not found"))
		} else {
			w.Write([]byte(resp))
		}
	})

	mux.HandleFunc("PUT /{table}/{key}/{value}", func(w http.ResponseWriter, r *http.Request) {
		table := r.PathValue("table")
		key := r.PathValue("key")
		value := r.PathValue("value")

		setVal(table, key, value, requestChannel)

		w.Write([]byte("OK"))
	})

	mux.HandleFunc("PUT /{table}", func(w http.ResponseWriter, r *http.Request) {
		table := r.PathValue("table")

		newTable(table, requestChannel)

		w.Write([]byte("OK"))
	})

	mux.HandleFunc("DELETE /{table}/{key}", func(w http.ResponseWriter, r *http.Request) {
		table := r.PathValue("table")
		key := r.PathValue("key")

		delKey(table, key, requestChannel)
	})

	mux.HandleFunc("DELETE /{table}", func(w http.ResponseWriter, r *http.Request) {
		table := r.PathValue("table")

		delTable(table, requestChannel)

		w.Write([]byte("OK"))
	})
}
