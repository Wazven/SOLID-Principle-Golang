package after

import "fmt"

func Log(pesan string){
	fmt.Println(pesan)
}

func LogBukuPinjam(pelangganID int, judulBuku string){
	fmt.Printf("Pelanggan dengan ID %d meminjam buku '%s'.\n", pelangganID, judulBuku)
}

func LogBukuKembali(judulBuku string){
	fmt.Printf("Buku '%s' telah dikembalikan.\n", judulBuku)
}

func LogPelangganPinjam(pelangganID int, judulBuku string){
	fmt.Printf("Pelanggan dengan ID %d meminjam buku '%s'.\n", pelangganID, judulBuku)
}

func LogPelangganKembali(pelangganID int, judulBuku string){
	fmt.Printf("Pelanggan dengan ID %d mengembalikan buku '%s'.\n", pelangganID, judulBuku)
}