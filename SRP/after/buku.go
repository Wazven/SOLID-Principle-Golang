package after

import "fmt"

type Buku struct {
	Judul    string
	Penulis  string
	Dipinjam bool //status buku
	PelangganID int //menyimpan ID pelanggan yang meminjam buku
}

func (b *Buku) Pinjam(pelangganID int) {
	if !b.Dipinjam {
		//status buku belum dipinjam
		b.Dipinjam = true
		b.PelangganID = pelangganID
		fmt.Printf("Buku '%s' dipinjam oleh pelanggan dengan ID %d. \n", b.Judul, pelangganID)
	} else {
		//buku sudah dipinjam
		fmt.Printf("Buku '%s' sudah dipinjam oleh pelanggan dengan ID %d. \n", b.Judul, pelangganID)
	}
}

func (b *Buku) Kembali() {
	//logika kembali
	if b.Dipinjam{
		//kembalikan buku
		b.Dipinjam = false
		fmt.Printf("Buku '%s' telah dikembalikan. \n", b.Judul)
	} else {
		//tidak dipinjam
		fmt.Printf("Buku '%s' tidak dapat dikembalikan karena belum dipinjam", b.Judul)
	}
}