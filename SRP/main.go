package main

import (
	"fmt"

	"github.com/Wazven/SOLID-Principle-Golang/SRP/after"
)

func main() {
	//buat intance buku
	buku := &after.Buku{
		Judul: "SRP Golang",
		Penulis: "Aldi",
	}

	//intance pelanggan
	// Membuat instance pelanggan
	pelanggan := &after.Pelanggan{}

	// Meminta pengguna memasukkan ID pelanggan dan nama
	var pelangganID int
	var pelangganName string

	fmt.Print("Masukkan ID Pelanggan: ")
	fmt.Scanln(&pelangganID)

	fmt.Print("Masukkan Nama Pelanggan: ")
	fmt.Scanln(&pelangganName)

	pelanggan.ID = pelangganID
	pelanggan.Name = pelangganName

	// Menggunakan fungsi-fungsi log yang ada di package after
	after.Log("Program dimulai")

	// Pelanggan meminjam buku
	pelanggan.PinjamBuku(buku)
	after.LogPelangganPinjam(pelanggan.ID, buku.Judul)

	// Pelanggan mengembalikan buku
	pelanggan.KembaliBuku(buku)
	after.LogPelangganKembali(pelanggan.ID, buku.Judul)

	after.Log("Program selesai")
}