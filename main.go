package main

import (
	"fmt"
	"os"
)

type Orang struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var orang = []Orang{
	{
		Nama:      "Agus",
		Alamat:    "Bandung",
		Pekerjaan: "Mahasiswa",
		Alasan:    "Mencari ilmu",
	},
	{
		Nama:      "Adi Wahyudi",
		Alamat:    "Semarang",
		Pekerjaan: "Mahasiswa",
		Alasan:    "Mempermudah mencari kerja",
	},
	{
		Nama:      "Irfan Maulana",
		Alamat:    "Ciamis",
		Pekerjaan: "Belum Bekerja",
		Alasan:    "Mencari sertifikasi",
	},
	{
		Nama:      "Aldi Dwi Prasetyo",
		Alamat:    "Bandung",
		Pekerjaan: "Mahasiswa",
		Alasan:    "Mencari sertifikasi",
	},
}

func parseAbsen(absen string) int {
	var hasil int
	fmt.Sscanf(absen, "%d", &hasil)
	return hasil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <absen>")
		return
	}
	absen := os.Args[1]

	ditemukan := false
	for i, orang := range orang {
		if i+1 == parseAbsen(absen) {
			ditemukan = true
			fmt.Printf("Nama     : %s\n", orang.Nama)
			fmt.Printf("Alamat   : %s\n", orang.Alamat)
			fmt.Printf("Pekerjaan: %s\n", orang.Pekerjaan)
			fmt.Printf("Alasan   : %s\n", orang.Alasan)
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan")
	}
}
