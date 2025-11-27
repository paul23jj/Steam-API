package routeur

import (
	"Steam-API/controller"
	"net/http"
)

func New() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	//a modifier pour mettre bien tout les chemins des pages des qu'on aura tt les pages
	//mux.HandleFunc("...", controller."")
}
