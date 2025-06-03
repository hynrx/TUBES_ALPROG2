package main

import (
	"fmt"
	"strings"
)

const jumlahMaksimumPengguna = 5
const jumlahMaksimumMinat = 10
const jumlahMaksimumKeahlian = 10
const jumlahMaksimumKarier = 10

type Pengguna struct {
	namaPengguna   string
	kataSandi      string
	daftarMinat    [jumlahMaksimumMinat]string
	daftarKeahlian [jumlahMaksimumKeahlian]string
	jumlahMinat    int
	jumlahKeahlian int
}

type Karier struct {
	namaKarier   string
	kategori     string
	gajiRataRata int
}

var daftarPengguna [jumlahMaksimumPengguna]Pengguna
var daftarKarier = [jumlahMaksimumKarier]Karier{
	{"Software Developer", "Pemrogramman", 12000000},
	{"Web Developer", "Pemrogramman", 5000000},
	{"Analisis Data", "DataScience", 5000000},
	{"Data Konsultan", "DataScience", 3500000},
	{"UI Desainer", "Desain", 4500000},
	{"Desain Grafis", "Desain", 3500000},
}
var jumlahKarier = 6

func main() {
	var penggunaSaatIni *Pengguna = nil
	var jumPengguna int = 0
	for {
		if penggunaSaatIni == nil {
			keluar := tampilkanMenuLogin(&jumPengguna, &penggunaSaatIni)
			if keluar {
				break
			}
		} else {
			menuUtama(penggunaSaatIni)
			penggunaSaatIni = nil
		}
	}
}

func tampilkanMenuLogin(currentJumlahPengguna *int, penggunaLogin **Pengguna) bool {
	fmt.Println("=====| APLIKASI REKOMENDASI KARIER |=====")
	fmt.Println("1. Daftar")
	fmt.Println("2. Masuk")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu: ")
	var pilihan string
	fmt.Scanln(&pilihan)
	switch pilihan {
	case "1":
		daftarPenggunaBaru(currentJumlahPengguna)
		return false
	case "2":
		loginPengguna(*currentJumlahPengguna, penggunaLogin)
		return false
	case "3":
		fmt.Println("=====| Selamat tinggal |======")
		fmt.Println("Program ini dibuat oleh:")
		fmt.Println("-Hisyam Nurdiatmoko")
		fmt.Println("-Mohammad Reyhan Aretha Fatin")
		*penggunaLogin = nil
		return true
	default:
		fmt.Println("Pilihan tidak valid.")
		return false
	}
}

func daftarPenggunaBaru(currentJumlahPengguna *int) {
	if *currentJumlahPengguna >= jumlahMaksimumPengguna {
		fmt.Println("Daftar penuh, tidak bisa tambah pengguna baru.")
		return
	}
	fmt.Print("Masukkan nama pengguna: ")
	namaPengguna := bacaInput()
	if cariIndeksPengguna(namaPengguna, *currentJumlahPengguna) != -1 {
		fmt.Println("Nama pengguna sudah digunakan.")
		return
	}
	fmt.Print("Masukkan kata sandi: ")
	kataSandi := bacaInput()
	daftarPengguna[*currentJumlahPengguna] = Pengguna{
		namaPengguna: namaPengguna,
		kataSandi:    kataSandi,
	}
	*currentJumlahPengguna++
	fmt.Println("Daftar berhasil, silakan masuk.")
}

func loginPengguna(currentJumlahPengguna int, penggunaLogin **Pengguna) {
	fmt.Print("Nama pengguna: ")
	namaPengguna := bacaInput()
	fmt.Print("Kata sandi: ")
	kataSandi := bacaInput()
	indeks := cariIndeksPengguna(namaPengguna, currentJumlahPengguna)
	if indeks == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}
	if daftarPengguna[indeks].kataSandi != kataSandi {
		fmt.Println("Kata sandi salah.")
		return
	}
	*penggunaLogin = &daftarPengguna[indeks]
	fmt.Printf("Selamat datang, %s!\n", (*penggunaLogin).namaPengguna)
}

func cariIndeksPengguna(namaPengguna string, currentJumlahPengguna int) int {
	for i := 0; i < currentJumlahPengguna; i++ {
		if daftarPengguna[i].namaPengguna == namaPengguna {
			return i
		}
	}
	return -1
}

func menuUtama(penggunaSaatIni *Pengguna) {
	for {
		fmt.Println("\n=====| MENU UTAMA |=====")
		var minatDisplay []string
		if penggunaSaatIni.jumlahMinat > 0 {
			minatDisplay = penggunaSaatIni.daftarMinat[0:penggunaSaatIni.jumlahMinat]
		}
		var keahlianDisplay []string
		if penggunaSaatIni.jumlahKeahlian > 0 {
			keahlianDisplay = penggunaSaatIni.daftarKeahlian[0:penggunaSaatIni.jumlahKeahlian]
		}
		fmt.Printf("Minat Anda: %v\n", minatDisplay)
		fmt.Printf("Keahlian Anda: %v\n", keahlianDisplay)
		fmt.Println("1. Kelola Minat")
		fmt.Println("2. Kelola Keahlian")
		fmt.Println("3. Lihat dan cari jalur karier")
		fmt.Println("4. Dapatkan rekomendasi karier")
		fmt.Println("5. Keluar (Logout)")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			kelolaMinat(penggunaSaatIni)
		case "2":
			kelolaKeahlian(penggunaSaatIni)
		case "3":
			cariKarier()
		case "4":
			menuRekomendasiKarier(penggunaSaatIni)
		case "5":
			fmt.Println("=====| Anda telah Logout |======")
			fmt.Println("Program ini dibuat oleh:")
			fmt.Println("-Hisyam Nurdiatmoko")
			fmt.Println("-Mohammad Reyhan Aretha Fatin")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func mengandungMinat(daftarMinat [jumlahMaksimumMinat]string, jumlahMinat int, nilai string) bool {
	for i := 0; i < jumlahMinat; i++ {
		if strings.EqualFold(daftarMinat[i], nilai) {
			return true
		}
	}
	return false
}

func cariIndeksMinat(daftarMinat [jumlahMaksimumMinat]string, jumlahMinat int, nilai string) int {
	for i := 0; i < jumlahMinat; i++ {
		if strings.EqualFold(daftarMinat[i], nilai) {
			return i
		}
	}
	return -1
}

func mengandungKeahlian(daftarKeahlian [jumlahMaksimumKeahlian]string, jumlahKeahlian int, nilai string) bool {
	for i := 0; i < jumlahKeahlian; i++ {
		if strings.EqualFold(daftarKeahlian[i], nilai) {
			return true
		}
	}
	return false
}

func cariIndeksKeahlian(daftarKeahlian [jumlahMaksimumKeahlian]string, jumlahKeahlian int, nilai string) int {
	for i := 0; i < jumlahKeahlian; i++ {
		if strings.EqualFold(daftarKeahlian[i], nilai) {
			return i
		}
	}
	return -1
}

func kelolaMinat(penggunaSaatIni *Pengguna) {
	for {
		fmt.Println("\n=====| Kelola Minat |=====")
		var minatDisplay []string
		if penggunaSaatIni.jumlahMinat > 0 {
			minatDisplay = penggunaSaatIni.daftarMinat[0:penggunaSaatIni.jumlahMinat]
		}
		fmt.Printf("Minat saat ini: %v\n", minatDisplay)
		fmt.Println("1. Tambah Minat")
		fmt.Println("2. Hapus Minat")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			if penggunaSaatIni.jumlahMinat >= jumlahMaksimumMinat {
				fmt.Println("Minat sudah penuh.")
			} else {
				fmt.Print("Masukkan minat baru: ")
				minatBaru := bacaInput()
				if !mengandungMinat(penggunaSaatIni.daftarMinat, penggunaSaatIni.jumlahMinat, minatBaru) {
					penggunaSaatIni.daftarMinat[penggunaSaatIni.jumlahMinat] = minatBaru
					penggunaSaatIni.jumlahMinat++
					fmt.Println("Minat berhasil ditambahkan.")
				} else {
					fmt.Println("Minat sudah ada.")
				}
			}
		case "2":
			if penggunaSaatIni.jumlahMinat == 0 {
				fmt.Println("Tidak ada minat untuk dihapus.")
			} else {
				fmt.Print("Masukkan minat yang ingin dihapus: ")
				minatHapus := bacaInput()
				indeks := cariIndeksMinat(penggunaSaatIni.daftarMinat, penggunaSaatIni.jumlahMinat, minatHapus)
				if indeks == -1 {
					fmt.Println("Minat tidak ditemukan.")
				} else {
					for i := indeks; i < penggunaSaatIni.jumlahMinat-1; i++ {
						penggunaSaatIni.daftarMinat[i] = penggunaSaatIni.daftarMinat[i+1]
					}
					penggunaSaatIni.daftarMinat[penggunaSaatIni.jumlahMinat-1] = ""
					penggunaSaatIni.jumlahMinat--
					fmt.Println("Minat berhasil dihapus.")
				}
			}
		case "3":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func kelolaKeahlian(penggunaSaatIni *Pengguna) {
	for {
		fmt.Println("\n=====| Kelola Keahlian |=====")
		var keahlianDisplay []string
		if penggunaSaatIni.jumlahKeahlian > 0 {
			keahlianDisplay = penggunaSaatIni.daftarKeahlian[0:penggunaSaatIni.jumlahKeahlian]
		}
		fmt.Printf("Keahlian saat ini: %v\n", keahlianDisplay)
		fmt.Println("1. Tambah Keahlian")
		fmt.Println("2. Hapus Keahlian")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			if penggunaSaatIni.jumlahKeahlian >= jumlahMaksimumKeahlian {
				fmt.Println("Keahlian sudah penuh.")
			} else {
				fmt.Print("Masukkan keahlian baru: ")
				keahlianBaru := bacaInput()
				if !mengandungKeahlian(penggunaSaatIni.daftarKeahlian, penggunaSaatIni.jumlahKeahlian, keahlianBaru) {
					penggunaSaatIni.daftarKeahlian[penggunaSaatIni.jumlahKeahlian] = keahlianBaru
					penggunaSaatIni.jumlahKeahlian++
					fmt.Println("Keahlian berhasil ditambahkan.")
				} else {
					fmt.Println("Keahlian sudah ada.")
				}
			}
		case "2":
			if penggunaSaatIni.jumlahKeahlian == 0 {
				fmt.Println("Tidak ada keahlian untuk dihapus.")
			} else {
				fmt.Print("Masukkan keahlian yang ingin dihapus: ")
				keahlianHapus := bacaInput()
				indeks := cariIndeksKeahlian(penggunaSaatIni.daftarKeahlian, penggunaSaatIni.jumlahKeahlian, keahlianHapus)
				if indeks == -1 {
					fmt.Println("Keahlian tidak ditemukan.")
				} else {
					for i := indeks; i < penggunaSaatIni.jumlahKeahlian-1; i++ {
						penggunaSaatIni.daftarKeahlian[i] = penggunaSaatIni.daftarKeahlian[i+1]
					}
					penggunaSaatIni.daftarKeahlian[penggunaSaatIni.jumlahKeahlian-1] = ""
					penggunaSaatIni.jumlahKeahlian--
					fmt.Println("Keahlian berhasil dihapus.")
				}
			}
		case "3":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func cariKarier() {
	fmt.Println("\n=====| Cari Jalur Karier |=====")
	fmt.Print("Cari dengan (1) Nama / (2) Kategori: ")
	var pilihan string
	fmt.Scanln(&pilihan)
	var kataKunci string
	switch pilihan {
	case "1":
		fmt.Print("Masukkan nama karier: ")
		kataKunci = strings.ToLower(bacaInput())
		fmt.Println("Hasil pencarian:")
		ditemukan := false
		for i := 0; i < jumlahKarier; i++ {
			if strings.Contains(strings.ToLower(daftarKarier[i].namaKarier), kataKunci) {
				fmt.Printf("- %s (Kategori: %s, Gaji: Rp %d)\n", daftarKarier[i].namaKarier, daftarKarier[i].kategori, daftarKarier[i].gajiRataRata)
				ditemukan = true
			}
		}
		if !ditemukan {
			fmt.Println("Tidak ditemukan karier dengan nama tersebut.")
		}
	case "2":
		fmt.Print("Masukkan kategori karier: ")
		kataKunci = strings.ToLower(bacaInput())
		fmt.Println("Hasil pencarian:")
		ditemukan := false
		for i := 0; i < jumlahKarier; i++ {
			if strings.ToLower(daftarKarier[i].kategori) == kataKunci {
				fmt.Printf("- %s (Kategori: %s, Gaji: Rp %d)\n", daftarKarier[i].namaKarier, daftarKarier[i].kategori, daftarKarier[i].gajiRataRata)
				ditemukan = true
			}
		}
		if !ditemukan {
			fmt.Println("Tidak ditemukan karier dengan kategori tersebut.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func menuRekomendasiKarier(penggunaSaatIni *Pengguna) {
	type Rekomendasi struct {
		karier      Karier
		persenCocok int
	}
	var daftarRekomendasi [jumlahMaksimumKarier]Rekomendasi
	jumlahRekomendasi := 0
	for i := 0; i < jumlahKarier; i++ {
		kategoriKarier := daftarKarier[i].kategori
		minatCocok := false
		if mengandungMinat(penggunaSaatIni.daftarMinat, penggunaSaatIni.jumlahMinat, kategoriKarier) {
			minatCocok = true
		}
		if !minatCocok {
			daftarRekomendasi[jumlahRekomendasi] = Rekomendasi{daftarKarier[i], 0}
			jumlahRekomendasi++
		} else {
			jumlahKeahlianCocok := 0
			for k := 0; k < penggunaSaatIni.jumlahKeahlian; k++ {
				if minatKeahlianKategori(penggunaSaatIni.daftarKeahlian[k], kategoriKarier) {
					jumlahKeahlianCocok++
				}
			}
			persenKecocokan := 0
			if jumlahKeahlianCocok == 1 {
				persenKecocokan = 50
			} else if jumlahKeahlianCocok >= 2 {
				persenKecocokan = 100
			}
			daftarRekomendasi[jumlahRekomendasi] = Rekomendasi{daftarKarier[i], persenKecocokan}
			jumlahRekomendasi++
		}
	}
	fmt.Println("\n=====| Pilih cara mengurutkan rekomendasi karier: |=====")
	fmt.Println("1. Berdasarkan persentase kecocokan")
	fmt.Println("2. Berdasarkan gaji rata-rata")
	fmt.Print("Pilih kriteria (Masukan angka): ")
	var pilihanKriteria string
	fmt.Scanln(&pilihanKriteria)
	fmt.Println("Pilih urutan:")
	fmt.Println("2. Menurun (Descending)")
	fmt.Println("2. Menaik (Ascending)")
	fmt.Print("Pilih urutan (Masukan angka): ")
	var pilihanUrutan string
	fmt.Scanln(&pilihanUrutan)
	for i := 0; i < jumlahRekomendasi-1; i++ {
		indeksPilihan := i
		for j := i + 1; j < jumlahRekomendasi; j++ {
			lebihBaik := false
			if pilihanKriteria == "1" {
				if pilihanUrutan == "1" {
					if daftarRekomendasi[j].persenCocok > daftarRekomendasi[indeksPilihan].persenCocok {
						lebihBaik = true
					}
				} else if pilihanUrutan == "2" {
					if daftarRekomendasi[j].persenCocok < daftarRekomendasi[indeksPilihan].persenCocok {
						lebihBaik = true
					}
				} else {
					fmt.Println("Pilihan urutan tidak valid, menggunakan Menurun.")
					if daftarRekomendasi[j].persenCocok > daftarRekomendasi[indeksPilihan].persenCocok {
						lebihBaik = true
					}
				}
			} else if pilihanKriteria == "2" {
				if pilihanUrutan == "1" {
					if daftarRekomendasi[j].karier.gajiRataRata > daftarRekomendasi[indeksPilihan].karier.gajiRataRata {
						lebihBaik = true
					}
				} else if pilihanUrutan == "2" {
					if daftarRekomendasi[j].karier.gajiRataRata < daftarRekomendasi[indeksPilihan].karier.gajiRataRata {
						lebihBaik = true
					}
				} else {
					fmt.Println("Pilihan urutan tidak valid, menggunakan Menurun.")
					if daftarRekomendasi[j].karier.gajiRataRata > daftarRekomendasi[indeksPilihan].karier.gajiRataRata {
						lebihBaik = true
					}
				}
			} else {
				fmt.Println("Pilihan kriteria tidak valid, mengurutkan berdasarkan persentase kecocokan (Menurun).")
				if daftarRekomendasi[j].persenCocok > daftarRekomendasi[indeksPilihan].persenCocok {
					lebihBaik = true
				}
				pilihanUrutan = "1"
			}

			if lebihBaik {
				indeksPilihan = j
			}
		}
		if indeksPilihan != i {
			sementara := daftarRekomendasi[i]
			daftarRekomendasi[i] = daftarRekomendasi[indeksPilihan]
			daftarRekomendasi[indeksPilihan] = sementara
		}
	}
	fmt.Println("\n=====| Rekomendasi Karier untuk Anda |=====")
	if jumlahRekomendasi == 0 {
		fmt.Println("Tidak ada rekomendasi yang dapat ditampilkan.")
	} else {
		for i := 0; i < jumlahRekomendasi; i++ {
			fmt.Printf("%d. %s (Kategori: %s) - Kecocokan: %d%% - Gaji Rata-rata: Rp %d\n",
				i+1,
				daftarRekomendasi[i].karier.namaKarier,
				daftarRekomendasi[i].karier.kategori,
				daftarRekomendasi[i].persenCocok,
				daftarRekomendasi[i].karier.gajiRataRata)
		}
	}
}

func minatKeahlianKategori(keahlian string, kategori string) bool {
	keahlianLower := strings.ToLower(keahlian)
	kategoriLower := strings.ToLower(kategori)
	switch kategoriLower {
	case "pemrogramman":
		if strings.Contains(keahlianLower, "python") || strings.Contains(keahlianLower, "go") || strings.Contains(keahlianLower, "java") {
			return true
		}
	case "desain":
		if strings.Contains(keahlianLower, "photoshop") || strings.Contains(keahlianLower, "figma") || strings.Contains(keahlianLower, "illustrator") {
			return true
		}
	case "datascience":
		if strings.Contains(keahlianLower, "sql") || strings.Contains(keahlianLower, "excel") || strings.Contains(keahlianLower, "data") {
			return true
		}
	}
	return false
}

func bacaInput() string {
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}
