package main
import "fmt"

func tampilkanData() {
    fmt.Println("👋 Keluar.")
    return
}

func tambahData() {
    fmt.Println("👋 Keluar.")
    return
}

func ubahData() {
    fmt.Println("👋 Keluar.")
    return
}

func hapusData() {
    fmt.Println("👋 Keluar.")
    return
}

func laporanBulanan() {
    fmt.Println("👋 Keluar.")
    return
}

func sequentialSearch(kategori string) {
    fmt.Println("👋 Keluar.")
    return
}

func binarySearch(kategori string) {
    fmt.Println("👋 Keluar.")
    return
}

func selectionSortPenghasilan(){
    
}

func insertionSortKategori(){
    
}

func menu() {
    var pilihan int = -1
    for {
        fmt.Println("\n=== MENU ===")
        fmt.Println("1. Tampilkan semua data")
        fmt.Println("2. Tambah data")
        fmt.Println("3. Ubah data")
        fmt.Println("4. Hapus data")
        fmt.Println("5. Laporan bulanan")
        fmt.Println("6. Cari kategori (Sequential)")
        fmt.Println("7. Cari kategori (Binary)")
        fmt.Println("8. Urutkan berdasarkan penghasilan")
        fmt.Println("9. Urutkan berdasarkan kategori")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&pilihan)
        switch pilihan {
        case 1:
            tampilkanData()
        case 2:
            tambahData()
        case 3:
            ubahData()
        case 4:
            hapusData()
        case 5:
            laporanBulanan()
        case 6:
            var kategori string
            fmt.Print("Kategori yang dicari: ")
            fmt.Scanln(&kategori)
            sequentialSearch(kategori)
        case 7:
            var kategori string
            fmt.Print("Kategori yang dicari: ")
            fmt.Scanln(&kategori)
            binarySearch(kategori)
        case 8:
            selectionSortPenghasilan()
            tampilkanData()
        case 9:
            insertionSortKategori()
            tampilkanData()
        case 0:
            fmt.Println("👋 Keluar.")
            return
        default:
            fmt.Println("❌ Pilihan tidak valid.")
        }
    }
}
func main(){
    menu()
}