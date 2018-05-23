package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//1.a.  set routing dari petugas
	router = setPetugasRouters(router)
	return router
}
