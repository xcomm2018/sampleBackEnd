package routers

import (
	"day14/controllers"

	"github.com/gorilla/mux"
)

func setPetugasRouters(router *mux.Router) *mux.Router {
	// 1.b buat fungsi  di controller  nya dlu
	router.HandleFunc("/petugas", controllers.GetPetugas).Methods("GET")
	return router
}
