package main

import (
	"fmt"
	"strings"
)

const jumlahMaksimumPengguna = 6
const jumlahMaksimumMinat = 6
const jumlahMaksimumKeahlian = 6
const jumlahMaksimumKarier = 6

// penggunaan array statis pada daftar minat dan daftar keahlian
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

// penggunaan array statis pada daftar pengguna dan daftar karier
var daftarPengguna [jumlahMaksimumPengguna]Pengguna
var jumlahPengguna int = 0
var penggunaSedangLogin *Pengguna = nil

var daftarKarier = [jumlahMaksimumKarier]Karier{
	{"Software Developer", "Pemrograman", 12000000},
	{"Web Developer", "Pemrograman", 5000000},
	{"Analisis Data", "DataScience", 5000000},
	{"Data Konsultan", "DataScience", 3500000},
	{"UI Desainer", "Desain", 4500000},
	{"Desain Grafis", "Desain", 3500000},
}
var jumlahKarier = 6

func main() {
	for {
		if penggunaSedangLogin == nil {
			menuMasuk()
		} else {
			menuUtama()
		}
	}
}

// fungsi menu masuk pengguna
func menuMasuk() {
	fmt.Println("=====| APLIKASI REKOMENDASI KARIER |=====")
	fmt.Println("1. Daftar")
	fmt.Println("2. Masuk")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu: ")
	var pilihan string
	fmt.Scanln(&pilihan)
	switch pilihan {
	case "1":
		daftarPenggunaBaru()
	case "2":
		loginPengguna()
	case "3":
		fmt.Println("=====| Selamat tinggal |======")
		fmt.Println("Program ini dibuat oleh:")
		fmt.Println("-Hisyam Nurdiatmoko")
		fmt.Println("-Mohammad Reyhan Aretha Fatin")
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// fungsi untuk pengguna baru yang mau daftar
func daftarPenggunaBaru() {
	if jumlahPengguna >= jumlahMaksimumPengguna {
		fmt.Println("Daftar penuh, tidak bisa tambah pengguna baru.")
		return
	}
	fmt.Print("Masukkan nama pengguna: ")
	namaPengguna := bacaInput()
	if cariIndeksPengguna(namaPengguna) != -1 {
		fmt.Println("Nama pengguna sudah digunakan.")
		return
	}
	fmt.Print("Masukkan kata sandi: ")
	kataSandi := bacaInput()
	daftarPengguna[jumlahPengguna] = Pengguna{
		namaPengguna: namaPengguna,
		kataSandi:    kataSandi,
	}
	jumlahPengguna++
	fmt.Println("Daftar berhasil, silakan masuk.")
}

// fungsi login pengguna
func loginPengguna() {
	fmt.Print("Nama pengguna: ")
	namaPengguna := bacaInput()
	fmt.Print("Kata sandi: ")
	kataSandi := bacaInput()
	indeks := cariIndeksPengguna(namaPengguna)
	if indeks == -1 {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}
	if daftarPengguna[indeks].kataSandi != kataSandi {
		fmt.Println("Kata sandi salah.")
		return
	}
	penggunaSedangLogin = &daftarPengguna[indeks]
	fmt.Printf("Selamat datang, %s!\n", penggunaSedangLogin.namaPengguna)
}

// fungsi cari indeks pengguna
func cariIndeksPengguna(namaPengguna string) int {
	for i := 0; i < jumlahPengguna; i++ {
		if strings.EqualFold(daftarPengguna[i].namaPengguna, namaPengguna) {
			return i
		}
	}
	return -1
}

// fungsi menampilkan menu utama
func menuUtama() {
	for {
		fmt.Println("\n=====| MENU UTAMA |=====")
		fmt.Printf("Minat Anda: %v\n", penggunaSedangLogin.daftarMinat[:penggunaSedangLogin.jumlahMinat])
		fmt.Printf("Keahlian Anda: %v\n", penggunaSedangLogin.daftarKeahlian[:penggunaSedangLogin.jumlahKeahlian])
		fmt.Println("1. Kelola Minat")
		fmt.Println("2. Kelola Keahlian")
		fmt.Println("3. Lihat dan cari jalur karier")
		fmt.Println("4. Dapatkan rekomendasi karier")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			kelolaMinat()
		case "2":
			kelolaKeahlian()
		case "3":
			cariKarier()
		case "4":
			menuRekomendasiKarier()
		case "5":
			fmt.Println("=====| Selamat tinggal |======")
			fmt.Println("Program ini dibuat oleh:")
			fmt.Println("-Hisyam Nurdiatmoko")
			fmt.Println("-Mohammad Reyhan Aretha Fatin")
			penggunaSedangLogin = nil
			return // kembali ke menu masuk
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// fungsi kelola minat
func kelolaMinat() {
	for {
		fmt.Println("\n=====| Kelola Minat |=====")
		fmt.Printf("Minat saat ini: %v\n", penggunaSedangLogin.daftarMinat[:penggunaSedangLogin.jumlahMinat])
		fmt.Println("1. Tambah Minat")
		fmt.Println("2. Hapus Minat")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			if penggunaSedangLogin.jumlahMinat >= jumlahMaksimumMinat {
				fmt.Println("Minat sudah penuh.")
			} else {
				fmt.Print("Masukkan minat baru: ")
				minatBaru := bacaInput()
				// penggunaan sequential search lewat fungsi mengandung
				if !mengandung(penggunaSedangLogin.daftarMinat[:penggunaSedangLogin.jumlahMinat], minatBaru) {
					penggunaSedangLogin.daftarMinat[penggunaSedangLogin.jumlahMinat] = minatBaru
					penggunaSedangLogin.jumlahMinat++
					fmt.Println("Minat berhasil ditambahkan.")
				} else {
					fmt.Println("Minat sudah ada.")
				}
			}
		case "2":
			if penggunaSedangLogin.jumlahMinat == 0 {
				fmt.Println("Tidak ada minat untuk dihapus.")
			} else {
				fmt.Print("Masukkan minat yang ingin dihapus: ")
				minatHapus := bacaInput()
				// penggunaan sequential search lewat fungsi cariIndeks
				indeks := cariIndeks(penggunaSedangLogin.daftarMinat[:penggunaSedangLogin.jumlahMinat], minatHapus)
				if indeks == -1 {
					fmt.Println("Minat tidak ditemukan.")
				} else {
					for i := indeks; i < penggunaSedangLogin.jumlahMinat-1; i++ {
						penggunaSedangLogin.daftarMinat[i] = penggunaSedangLogin.daftarMinat[i+1]
					}
					penggunaSedangLogin.jumlahMinat--
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

// fungsi kelola keahlian
func kelolaKeahlian() {
	for {
		fmt.Println("\n=====| Kelola Keahlian |=====")
		fmt.Printf("Keahlian saat ini: %v\n", penggunaSedangLogin.daftarKeahlian[:penggunaSedangLogin.jumlahKeahlian])
		fmt.Println("1. Tambah Keahlian")
		fmt.Println("2. Hapus Keahlian")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			if penggunaSedangLogin.jumlahKeahlian >= jumlahMaksimumKeahlian {
				fmt.Println("Keahlian sudah penuh.")
			} else {
				fmt.Print("Masukkan keahlian baru: ")
				keahlianBaru := bacaInput()
				// penggunaan sequential search lewat fungsi mengandung
				if !mengandung(penggunaSedangLogin.daftarKeahlian[:penggunaSedangLogin.jumlahKeahlian], keahlianBaru) {
					penggunaSedangLogin.daftarKeahlian[penggunaSedangLogin.jumlahKeahlian] = keahlianBaru
					penggunaSedangLogin.jumlahKeahlian++
					fmt.Println("Keahlian berhasil ditambahkan.")
				} else {
					fmt.Println("Keahlian sudah ada.")
				}
			}
		case "2":
			if penggunaSedangLogin.jumlahKeahlian == 0 {
				fmt.Println("Tidak ada keahlian untuk dihapus.")
			} else {
				fmt.Print("Masukkan keahlian yang ingin dihapus: ")
				keahlianHapus := bacaInput()
				// penggunaan sequential search lewat fungsi cariIndeks
				indeks := cariIndeks(penggunaSedangLogin.daftarKeahlian[:penggunaSedangLogin.jumlahKeahlian], keahlianHapus)
				if indeks == -1 {
					fmt.Println("Keahlian tidak ditemukan.")
				} else {
					for i := indeks; i < penggunaSedangLogin.jumlahKeahlian-1; i++ {
						penggunaSedangLogin.daftarKeahlian[i] = penggunaSedangLogin.daftarKeahlian[i+1]
					}
					penggunaSedangLogin.jumlahKeahlian--
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

// fungsi pembantu memeriksa apakah array mengandung nilai
func mengandung(arr []string, nilai string) bool {
	for _, v := range arr {
		if strings.EqualFold(v, nilai) {
			return true
		}
	}
	return false
}

// fungsi bantu untuk mencari indeks nilai dalam array
func cariIndeks(arr []string, nilai string) int {
	for i, v := range arr {
		if strings.EqualFold(v, nilai) {
			return i
		}
	}
	return -1
}

// fungsi cari karier
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
		// sequential search untuk mencari karier berdasarkan nama
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
		// sequential search untuk mencari karier berdasarkan kategori
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

// fungsi menu rekomendasi karir
func menuRekomendasiKarier() {
	type Rekomendasi struct {
		karier      Karier
		persenCocok int
	}
	var daftarRekomendasi [jumlahMaksimumKarier]Rekomendasi
	jumlahRekomendasi := 0
	for i := 0; i < jumlahKarier; i++ {
		kategoriKarier := daftarKarier[i].kategori
		minatCocok := false
		for j := 0; j < penggunaSedangLogin.jumlahMinat; j++ {
			if strings.EqualFold(penggunaSedangLogin.daftarMinat[j], kategoriKarier) {
				minatCocok = true
			}
		}
		if !minatCocok {
			daftarRekomendasi[jumlahRekomendasi] = Rekomendasi{daftarKarier[i], 0}
			jumlahRekomendasi++
		} else {
			jumlahKeahlianCocok := 0
			for k := 0; k < penggunaSedangLogin.jumlahKeahlian; k++ {
				if minatKeahlianKategori(penggunaSedangLogin.daftarKeahlian[k], kategoriKarier) {
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
	fmt.Println("1. Berdasarkan persentase kecocokan (turun)")
	fmt.Println("2. Berdasarkan persentase kecocokan (naik)")
	fmt.Println("3. Berdasarkan rata-rata gaji (turun)")
	fmt.Println("4. Berdasarkan rata-rata gaji (naik)")
	fmt.Print("Pilih (Masukan angka): ")
	var pilihan string
	fmt.Scanln(&pilihan)
	// proses pengurutan menggunakan selection sort
	switch pilihan {
	case "1": // urut persentase turun (descending)
		for i := 0; i < jumlahRekomendasi-1; i++ {
			indeksMaks := i
			for j := i + 1; j < jumlahRekomendasi; j++ {
				if daftarRekomendasi[j].persenCocok > daftarRekomendasi[indeksMaks].persenCocok {
					indeksMaks = j
				}
			}
			if indeksMaks != i {
				sementara := daftarRekomendasi[i]
				daftarRekomendasi[i] = daftarRekomendasi[indeksMaks]
				daftarRekomendasi[indeksMaks] = sementara
			}
		}
	case "2": // urut persentase naik (ascending)
		for i := 0; i < jumlahRekomendasi-1; i++ {
			indeksMin := i
			for j := i + 1; j < jumlahRekomendasi; j++ {
				if daftarRekomendasi[j].persenCocok < daftarRekomendasi[indeksMin].persenCocok {
					indeksMin = j
				}
			}
			if indeksMin != i {
				sementara := daftarRekomendasi[i]
				daftarRekomendasi[i] = daftarRekomendasi[indeksMin]
				daftarRekomendasi[indeksMin] = sementara
			}
		}
	case "3": // urut gaji turun (descending)
		for i := 0; i < jumlahRekomendasi-1; i++ {
			indeksMaks := i
			for j := i + 1; j < jumlahRekomendasi; j++ {
				if daftarRekomendasi[j].karier.gajiRataRata > daftarRekomendasi[indeksMaks].karier.gajiRataRata {
					indeksMaks = j
				}
			}
			if indeksMaks != i {
				sementara := daftarRekomendasi[i]
				daftarRekomendasi[i] = daftarRekomendasi[indeksMaks]
				daftarRekomendasi[indeksMaks] = sementara
			}
		}
	case "4": // urut gaji naik (ascending)
		for i := 0; i < jumlahRekomendasi-1; i++ {
			indeksMin := i
			for j := i + 1; j < jumlahRekomendasi; j++ {
				if daftarRekomendasi[j].karier.gajiRataRata < daftarRekomendasi[indeksMin].karier.gajiRataRata {
					indeksMin = j
				}
			}
			if indeksMin != i {
				sementara := daftarRekomendasi[i]
				daftarRekomendasi[i] = daftarRekomendasi[indeksMin]
				daftarRekomendasi[indeksMin] = sementara
			}
		}
	default:
		fmt.Println("Pilihan tidak valid. Mengurutkan berdasarkan persentase kecocokan menurun.")
		// urut persentase turun (descending)
		for i := 0; i < jumlahRekomendasi-1; i++ {
			indeksMaks := i
			for j := i + 1; j < jumlahRekomendasi; j++ {
				if daftarRekomendasi[j].persenCocok > daftarRekomendasi[indeksMaks].persenCocok {
					indeksMaks = j
				}
			}
			if indeksMaks != i {
				sementara := daftarRekomendasi[i]
				daftarRekomendasi[i] = daftarRekomendasi[indeksMaks]
				daftarRekomendasi[indeksMaks] = sementara
			}
		}
	}
	// fungsi menampilkan hasil rekomendasi yang sudah diurutkan
	fmt.Println("\n=====| Rekomendasi Karier untuk Anda |=====")
	for i := 0; i < jumlahRekomendasi; i++ {
		fmt.Printf("%d. %s (Kategori: %s) - Kecocokan: %d%% - Gaji Rata-rata: Rp %d\n",
			i+1,
			daftarRekomendasi[i].karier.namaKarier,
			daftarRekomendasi[i].karier.kategori,
			daftarRekomendasi[i].persenCocok,
			daftarRekomendasi[i].karier.gajiRataRata)
	}
}

// fungsi untuk memeriksa input keahlian
func minatKeahlianKategori(keahlian string, kategori string) bool {
	keahlianLower := strings.ToLower(keahlian)
	kategoriLower := strings.ToLower(kategori)
	switch kategoriLower {
	case "pemrograman":
		if strings.Contains(keahlianLower, "python") || strings.Contains(keahlianLower, "go") {
			return true
		}
	case "desain":
		if strings.Contains(keahlianLower, "photoshop") || strings.Contains(keahlianLower, "figma") {
			return true
		}
	case "datascience":
		if strings.Contains(keahlianLower, "sql") || strings.Contains(keahlianLower, "excel") {
			return true
		}
	}
	return false
}

// fungsi baca sampe baris baru
func bacaInput() string {
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}
