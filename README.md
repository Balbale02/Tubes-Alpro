# Aplikasi  Pelacak Side Hustle dan Passive Income
Aplikasi berbasis bahasa pemrograman Go (Golang) untuk mengelola berbagai sumber penghasilan yaitu : pekerjaan sampingan, freelancing, bisnis digital, atau investasi pasif, secara sederhana.
Aplikasi ini memungkinkan pengguna untuk menambah, mengubah, menghapus, mencari, mengurutkan, dan menganalisis data penghasilan dari berbagai kategori.
# Fitur
- Menambah sumber penghasilan dengan nama, kategori, penghasilan, dan biaya operasional
- Mengubah dan menghapus data penghasilan
- Mencari data berdasarkan kategori menggunakan Sequential Search dan Binary Search
- Mengurutkan data berdasarkan penghasilan (Selection Sort) dan kategori (Insertion Sort)
- Menampilkan laporan bulanan beserta rekomendasi keuangan
- Menampilkan semua data yang sudah tersimpan
 # Struktur Data
Menggunakan tipe data struct IncomeSource dengan atribut:
- Nama (string)
- Kategori (string)
- Penghasilan (float64)
- BiayaOperasional (float64)
Data disimpan dalam array statis dengan kapasitas maksimum 1000 elemen.
# Cara Menjalankan
1. Pastikan sudah menginstal Go di komputer Anda.
2. Simpan kode program dalam file `main.go` (atau `Tubes.Go`).
3. Jalankan perintah berikut di terminal:
	`go run main.go`
4. Ikuti menu yang muncul untuk mengoperasikan aplikasi.
# Menu Utama
```
1. Tampilkan semua data
2. Tambah data
3. Ubah data
4. Hapus data
5. Laporan bulanan
6. Cari kategori (Sequential)
7. Cari kategori (Binary)
8. Urutkan penghasilan (Asc/Desc)
9. Urutkan kategori (Asc/Desc)
10. Bersihkan layar
0. Keluar
```
# Contoh Input
```
Nama sumber: Toko Online
Pilih kategori:
1. Pekerjaan Sampingan
2. Freelancing
3. Bisnis Digital
4. Investasi Pasif
Masukkan nomor kategori: 3
Penghasilan: 5000000
Biaya operasional: 1500000
```
