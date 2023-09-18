package after

import "fmt"

type Pelanggan struct {
	ID   int
	Name string
}

func (p *Pelanggan) PinjamBuku(buku *Buku) {
	// Logika untuk meminjam buku oleh pelanggan
	if buku.Dipinjam {
		fmt.Printf("Pelanggan %s (ID: %d) tidak dapat meminjam buku '%s' karena buku sudah dipinjam.\n", p.Name, p.ID, buku.Judul)
	} else {
		buku.Pinjam(p.ID)
	}
}

func (p *Pelanggan) KembaliBuku(buku *Buku) {
	if buku.Dipinjam && buku.Judul != ""{
		//buku sedang dipinjam oleh pelanggan yang sesuai
		if p.ID == buku.PelangganID{
			buku.Kembali()
		} else {
			//buku dipinjam oleh pelanggan lain
			fmt.Printf("Pelanggan %s (ID: %d) tidak dapat mengembalikan buku '%s' karena buku dipinjam pelanggan lain. \n", p.Name, p.ID, buku.Judul)
		}
	} else {
		fmt.Printf("Pelanggan %s (ID: %d) tidak dapat mengembalikan buku '%s' karena buku tidak valid atau tidak sedang dipinjam. \n", p.Name, p.ID, buku.Judul)
	}
}