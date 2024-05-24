package main

import(
	"net/http"
)

func Routes() *http.ServeMux {
	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/endpoint01", endpoint01)
	multiplexer.HandleFunc("/endpoint02", endpoint02)
	multiplexer.HandleFunc("/endpoint03", endpoint03)

	return multiplexer
}