package before

import "fmt"

type Buku struct {
	Judul   string
	Penulis string
	Dipinjam bool //status buku
	PelangganID int //menyimpan ID pelanggan yang meminjam buku
}

type Pelanggan struct {
	ID   int
	Name string
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

func Log(pesan string) {
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


