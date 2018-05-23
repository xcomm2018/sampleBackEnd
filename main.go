package main

import (
	"day14/routers"
	"log"
	"net/http"
)

func main() {
	// p1. buat routing
	// buat fungsi routing
	router := routers.InitRouters()

	// buat configurasi server
	log.Fatal(http.ListenAndServe(":8887", router))

}
