package controllers

import (
	"day14/data"
	"encoding/json"
	"net/http"
)

func GetPetugas(w http.ResponseWriter, r *http.Request) {
	// ambil datanya
	// untuk ambil data maka perlu koneksi
	// 1.c. buat koneksi
	context := Context{}
	//defer context.Close()
	d := DBInitial(context.DB)
	defer d.Close()
	// ambil data dari repositori
	// buat perintah di repositori
	// 1.d  buat repo petugas
	repo := &data.PetugasRepository{d}

	petugas := data.GetAll(repo)
	// olah errornya
	// tampilin datanya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(petugas)
	w.Write(mdl)
}
