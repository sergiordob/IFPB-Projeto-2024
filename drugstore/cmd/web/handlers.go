package main

import (
	"fmt"
	"net/http"
)

func endpoint01(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Testando endpoint 01")
}

func endpoint02(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Testando endpoint 02")
}

func endpoint03(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Testando endpoint 03")
}
