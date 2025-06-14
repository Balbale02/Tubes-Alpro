package main

import "fmt"

func main() {
	fmt.Println("\nSelamat datang di Aplikasi  Pelacak Side Hustle dan Passive Income")
	menu()
}

type IncomeSource struct {
	Nama             string
	Kategori         string
	Penghasilan      float64
	BiayaOperasional float64
}

const Maks = 1000

var sumber [Maks]IncomeSource
var n int // jumlah data

var daftarKategori = []string{
	"Kerja Sampingan",
	"Freelancing",
	"Bisnis Digital",
	"Investasi Pasif",
}

func tampilKategori() {
	var i int
	for i = 0; i < len(daftarKategori); i++ {
		fmt.Printf("%d. %s\n", i+1, daftarKategori[i])
	}
}

func kategoriValid(input string) bool {
	var i int
	for i = 0; i < len(daftarKategori); i++ {
		if input == daftarKategori[i] {
			return true
		}
	}
	return false
}

func tampilkanData() {
	var i int = 0
	var laba float64
	fmt.Println("\n================================= Daftar Penghasilan =================================")
	fmt.Printf("%-4s | %-20s | %-18s | %-15s | %-15s\n", "No", "Nama", "Kategori", "Penghasilan", "Laba")
	for i = 0; i < n; i++ {
		laba = sumber[i].Penghasilan - sumber[i].BiayaOperasional
		fmt.Printf("[%d]  | %-20s | %-18s | Rp%-13.2f | Rp%-13.2f\n", i, sumber[i].Nama, sumber[i].Kategori, sumber[i].Penghasilan, laba)
	}
	// Jika tidak ada data
	if n == 0 {
		fmt.Println("		 	  ❌ Tidak ada data yang tersedia.")
	}
	fmt.Println("=======================================================================================\n")
}

func tambahData() {
	if n < Maks {
		var nama string
		var penghasilan, biaya float64
		var pilih int

		fmt.Print("\nNama sumber: ")
		fmt.Scanln(&nama)

		fmt.Println("Pilih kategori:")
		tampilKategori()
		fmt.Print("Masukkan nomor kategori: ")
		fmt.Scanln(&pilih)
		if pilih >= 1 && pilih <= len(daftarKategori) {
			fmt.Print("Penghasilan: ")
			fmt.Scanln(&penghasilan)
			fmt.Print("Biaya operasional: ")
			fmt.Scanln(&biaya)
			sumber[n] = IncomeSource{nama, daftarKategori[pilih-1], penghasilan, biaya}
			n++
			fmt.Println("✅ Data berhasil ditambahkan.")
		} else {
			fmt.Println("❌ Pilihan kategori tidak valid.\n")
			fmt.Println("Silakan pilih kategori yang valid dari daftar berikut:")
			tampilKategori()
		}
	} else {
		fmt.Println("❌ Kapasitas penuh.")
	}
}

func ubahData() {
	tampilkanData()
	var nama string
	var penghasilan, biaya float64
	var pilih int
	var index int
	fmt.Print("Masukkan indeks data yang ingin diubah: ")
	fmt.Scanln(&index)
	if index >= 0 && index < n {
		fmt.Print("Nama sumber baru: ")
		fmt.Scanln(&nama)
		fmt.Println("Pilih kategori baru:")
		tampilKategori()
		fmt.Print("Masukkan nomor kategori: ")
		fmt.Scanln(&pilih)
		if pilih >= 1 && pilih <= len(daftarKategori) {
			fmt.Print("Penghasilan baru: ")
			fmt.Scanln(&penghasilan)
			fmt.Print("Biaya operasional baru: ")
			fmt.Scanln(&biaya)
			sumber[index] = IncomeSource{nama, daftarKategori[pilih-1], penghasilan, biaya}
			fmt.Println("✅ Data berhasil diubah.")
		} else {
			fmt.Println("❌ Kategori tidak valid.\n")
		}

	} else {
		fmt.Println("❌ Indeks tidak valid.")
	}
}

func hapusData() {
	tampilkanData()
	var index int
	var i int
	fmt.Print("Masukkan indeks data yang ingin dihapus: ")
	fmt.Scanln(&index)
	if index >= 0 && index < n {
		i = index
		for i < n-1 {
			sumber[i] = sumber[i+1]
			i++
		}
		n = n - 1
		fmt.Println("✅ Data berhasil dihapus.")
	} else {
		fmt.Println("❌ Indeks tidak valid.\n")
	}
}

func laporanBulanan() {
	var total, laba float64
	var i int = 0
	for i < n {
		total = total + sumber[i].Penghasilan
		laba = laba + (sumber[i].Penghasilan - sumber[i].BiayaOperasional)
		i++
	}
	fmt.Println("\n=== Laporan Bulanan ===")
	fmt.Printf("Total Penghasilan: Rp%.2f\n", total)
	fmt.Printf("Total Laba: Rp%.2f\n", laba)
	if laba < 0.3*total {
		fmt.Println("Rekomendasi: Kurangi biaya operasional.")
	} else {
		fmt.Println("Rekomendasi: Laba sehat, pertahankan atau cari peluang baru.")
	}
}

func sequentialSearch(kategori string) {
	fmt.Println("\nHasil Pencarian (Sequential):")
	var i int = 0
	var ditemukan bool = false
	for i < n {
		if sumber[i].Kategori == kategori {
			fmt.Printf("%s | Rp%.2f\n", sumber[i].Nama, sumber[i].Penghasilan)
			fmt.Printf("Kategori ditemukan pada indeks %d\n", i)
			ditemukan = true
		}
		i++
	}
	if !ditemukan {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

func binarySearch(kategori string) {
	insertionSortKategori(true)
	var mid int
	var ditemukan bool = false
	var low, high int = 0, n - 1
	for low <= high && !ditemukan {
		mid = (low + high) / 2
		if sumber[mid].Kategori == kategori {
			ditemukan = true
		} else {
			if kategori < sumber[mid].Kategori {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	if ditemukan {
		fmt.Printf("%s ditemukan pada indeks %d\n", kategori, mid)
		fmt.Println("Hasil Pencarian (Binary):")
		fmt.Printf("%s | Rp%.2f\n", sumber[mid].Nama, sumber[mid].Penghasilan)
	} else {
		fmt.Println("Kategori tidak ditemukan (binary search).")
	}
}

func selectionSortPenghasilan(ascending bool) {
	var i int = 0
	var minIdx int
	var temp IncomeSource
	var j int
	for i < n-1 {
		minIdx = i
		j = i + 1
		for j < n {
			if (ascending && sumber[j].Penghasilan < sumber[minIdx].Penghasilan) || (!ascending && sumber[j].Penghasilan > sumber[minIdx].Penghasilan) {
				minIdx = j
			}
			j++
		}
		temp = sumber[i]
		sumber[i] = sumber[minIdx]
		sumber[minIdx] = temp
		i++
	}
}

func insertionSortKategori(ascending bool) {
	var i int = 1
	var j int
	var key IncomeSource
	for i < n {
		key = sumber[i]
		j = i - 1
		for j >= 0 && ((ascending && sumber[j].Kategori > key.Kategori) || (!ascending && sumber[j].Kategori < key.Kategori)) {
			sumber[j+1] = sumber[j]
			j--
		}
		sumber[j+1] = key
		i++
	}
}

func menu() {
	var pilihan int = -1
	for pilihan != 0 {
		// Menampilkan menu
		fmt.Println("\n   Menu Aplikasi Manajemen Penghasilan")
		fmt.Println("=========================================")
		fmt.Println("1. Tampilkan semua data")
		fmt.Println("2. Tambah data")
		fmt.Println("3. Ubah data")
		fmt.Println("4. Hapus data")
		fmt.Println("5. Laporan bulanan")
		fmt.Println("6. Cari kategori (Sequential)")
		fmt.Println("7. Cari kategori (Binary)")
		fmt.Println("8. Urutkan penghasilan (Asc/Desc)")
		fmt.Println("9. Urutkan kategori (Asc/Desc)")
		fmt.Println("10. Bersihkan layar")
		fmt.Println("0. Keluar")
		fmt.Println("=========================================")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			tampilkanData()
		} else if pilihan == 2 {
			tambahData()
		} else if pilihan == 3 {
			ubahData()
		} else if pilihan == 4 {
			hapusData()
		} else if pilihan == 5 {
			laporanBulanan()
		} else if pilihan == 6 {
			var kategori string
			fmt.Print("Masukkan kategori: ")
			tampilKategori()
			fmt.Scanln(&kategori)
			if kategoriValid(kategori) {
				sequentialSearch(kategori)
			} else {
				fmt.Println("❌ Kategori tidak valid. Gunakan salah satu dari:")
				tampilKategori()
			}
		} else if pilihan == 7 {
			var kategori string
			fmt.Print("Masukkan kategori: ")
			tampilKategori()
			fmt.Scanln(&kategori)
			if kategoriValid(kategori) {
				binarySearch(kategori)
			} else {
				fmt.Println("❌ Kategori tidak valid. Gunakan salah satu dari:")
				tampilKategori()
			}
		} else if pilihan == 8 {
			var mode string
			fmt.Print("Urutkan Ascending (a) atau Descending (d): ")
			fmt.Scanln(&mode)
			selectionSortPenghasilan(mode == "a")
			tampilkanData()
		} else if pilihan == 9 {
			var mode string
			fmt.Print("Urutkan Ascending (a) atau Descending (d): ")
			fmt.Scanln(&mode)
			insertionSortKategori(mode == "a")
			tampilkanData()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			return
		} else if pilihan == 10 {
			ClearScreen()
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
