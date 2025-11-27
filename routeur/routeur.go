package routeur

import (
	"Steam-API/controller"
	"net/http"
)

func New() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/aPropos", controller.aPropos)
	mux.HandleFunc("/categories", controller.categories)
	mux.HandleFunc("/collection", controller.collection)
	mux.HandleFunc("/favoris", controller.favoris)
	mux.HandleFunc("/recherche", controller.recherche)
	mux.HandleFunc("/resssources", controller.resssources)
	//normalement ya tt les pages

	return mux
}
