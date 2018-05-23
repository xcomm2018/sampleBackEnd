package data

import (
	"database/sql"
	"day14/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type PetugasRepository struct {
	DB *sql.DB
}

// untuk nilai return  get all butuh stuctur dari petugas
//2.a. buat model dari Petugas
func GetAll(db *PetugasRepository) []models.Petugas {
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select Nama From Petugas")

	if err != nil {
		fmt.Println("errorQuery")
		fmt.Println(err)
		return nil

	}

	defer result.Close()
	fmt.Println(result)
	petugas := []models.Petugas{}
	for result.Next() {
		var p models.Petugas
		if err := result.Scan(&p.Nama); err != nil {
			fmt.Println(err)
			return nil
		}
		petugas = append(petugas, p)

	}
	return petugas
}
