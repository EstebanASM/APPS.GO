package handler

import (
	"fmt"
	"net/http"
)

// Esta es la función exportada que Vercel invocará
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola soy Esteban y este es el deber de programación distribuida con el lenguaje Go")
}
