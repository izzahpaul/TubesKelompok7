package main

import (
	"fmt"
	"strings"
)

// Struct untuk menyimpan data penghasilan
type Penghasilan struct {
	NamaSumber   string
	Kategori     string
	Pendapatan   float64
	BiayaOperasi float64
}

// Slice untuk menyimpan data penghasilan
var data []Penghasilan

func main() {
	for {
		// Menu utama
		fmt.Println("=== APLIKASI PELACAK SIDE HUSTLE DAN PASSIVE INCOME ===")
		fmt.Println("[1] Tambah Sumber Pendapatan")
		fmt.Println("[2] Lihat Semua Penghasilan")
		fmt.Println("[3] Edit Penghasilan")
		fmt.Println("[4] Hapus Penghasilan")
		fmt.Println("[5] Cari Sumber Penghasilan (Sequential/Binary Search)")
		fmt.Println("[6] Urutkan Penghasilan (Selection/Insertion Sort)")
		fmt.Println("[7] Tampilkan Laporan Bulanan")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih menu> ")

		var pilihan int
		fmt.Scan(&pilihan)

		// Pemilihan menu berdasarkan input user
		switch pilihan {
		case 1:
			tambahPenghasilan()
		case 2:
			lihatPenghasilan()
		case 3:
			editPenghasilan()
		case 4:
			hapusPenghasilan()
		case 5:
			cariSumberPenghasilan()
		case 6:
			urutkanPenghasilan()
		case 7:
			laporanBulanan()
		case 0:
			fmt.Println("Terima kasih sudah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Println()
	}
}

// Menambahkan data penghasilan
func tambahPenghasilan() {
	var ph Penghasilan
	fmt.Print("Nama Sumber: ")
	fmt.Scan(&ph.NamaSumber)
	fmt.Print("Kategori: ")
	fmt.Scan(&ph.Kategori)
	fmt.Print("Pendapatan(Rp): ")
	fmt.Scan(&ph.Pendapatan)
	fmt.Print("Biaya Operasional(Rp): ")
	fmt.Scan(&ph.BiayaOperasi)
	data = append(data, ph)
	fmt.Println("Data berhasil ditambahkan!")
}

// Menampilkan semua data penghasilan
func lihatPenghasilan() {
	if len(data) == 0 {
		fmt.Println("Belum ada data yang ditambahkan.")
		return
	}
	fmt.Println("\n=== DAFTAR PENGHASILAN ===")
	for i, ph := range data {
		fmt.Printf("[%d] %s | Kategori: %s | Pendapatan: Rp%.0f | Laba: Rp%.0f\n", i+1, ph.NamaSumber, ph.Kategori, ph.Pendapatan, ph.Pendapatan-ph.BiayaOperasi)
	}
}

// Mengedit data penghasilan berdasarkan indeks
func editPenghasilan() {
	if len(data) == 0 {
		fmt.Println("Belum ada data untuk diedit.")
		return
	}
	lihatPenghasilan()
	var idx int
	fmt.Print("Nomor data yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx >= 0 && idx < len(data) {
		fmt.Print("Nama Sumber: ")
		fmt.Scan(&data[idx].NamaSumber)
		fmt.Print("Kategori: ")
		fmt.Scan(&data[idx].Kategori)
		fmt.Print("Pendapatan(Rp): ")
		fmt.Scan(&data[idx].Pendapatan)
		fmt.Print("Biaya Operasional(Rp): ")
		fmt.Scan(&data[idx].BiayaOperasi)
		fmt.Println("Data berhasil diperbaharui!")
	} else {
		fmt.Println("Nomor tidak valid.")
	}
}

// Menghapus data penghasilan berdasarkan indeks
func hapusPenghasilan() {
	if len(data) == 0 {
		fmt.Println("Belum ada data untuk dihapus.")
		return
	}
	lihatPenghasilan()
	var idx int
	fmt.Print("Nomor data yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx--
	if idx >= 0 && idx < len(data) {
		data = append(data[:idx], data[idx+1:]...)
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Nomor tidak valid.")
	}
}

// Pencarian menggunakan Sequential Search berdasarkan kategori
func SequentialSearch(data []Penghasilan, keyword string) []Penghasilan {
	var result []Penghasilan
	for _, d := range data {
		if strings.EqualFold(d.Kategori, keyword) {
			result = append(result, d)
		}
	}
	return result
}

// Pencarian menggunakan Binary Search (data harus sudah diurutkan berdasarkan kategori)
func BinarySearch(data []Penghasilan, keyword string) *Penghasilan {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		midKategori := strings.ToLower(data[mid].Kategori)

		if midKategori == strings.ToLower(keyword) {
			return &data[mid]
		} else if midKategori < strings.ToLower(keyword) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

// Mengurutkan data berdasarkan pendapatan (descending)
func SelectionSortByJumlah(data []Penghasilan) []Penghasilan {
	sorted := make([]Penghasilan, len(data))
	copy(sorted, data)

	for i := 0; i < len(sorted); i++ {
		maxIdx := i
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].Pendapatan > sorted[maxIdx].Pendapatan {
				maxIdx = j
			}
		}
		sorted[i], sorted[maxIdx] = sorted[maxIdx], sorted[i]
	}
	return sorted
}

// Mengurutkan data berdasarkan kategori (ascending)
func InsertionSortByKategori(data []Penghasilan) []Penghasilan {
	sorted := make([]Penghasilan, len(data))
	copy(sorted, data)

	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && strings.ToLower(sorted[j].Kategori) > strings.ToLower(key.Kategori) {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	return sorted
}

// Menu untuk pencarian penghasilan berdasarkan kategori
func cariSumberPenghasilan() {
	var pilihan int
	fmt.Println("\n--- Pencarian Berdasarkan Kategori ---")
	fmt.Println("[1] Sequential Search")
	fmt.Println("[2] Binary Search")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&pilihan)

	fmt.Print("Masukkan kategori: ")
	var kategori string
	fmt.Scan(&kategori)

	switch pilihan {
	case 1:
		result := SequentialSearch(data, kategori)
		for _, r := range result {
			fmt.Printf("%s | Kategori: %s | Pendapatan: Rp%.0f\n", r.NamaSumber, r.Kategori, r.Pendapatan)
		}
	case 2:
		sorted := InsertionSortByKategori(data)
		result := BinarySearch(sorted, kategori)
		if result != nil {
			fmt.Printf("%s | Kategori: %s | Pendapatan: Rp%.0f\n", result.NamaSumber, result.Kategori, result.Pendapatan)
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Menu untuk pengurutan data penghasilan
func urutkanPenghasilan() {
	var pilihan int
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("[1] Pendapatan (Selection Sort)")
	fmt.Println("[2] Kategori (Insertion Sort)")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		sorted := SelectionSortByJumlah(data)
		fmt.Println("Diurutkan berdasarkan pendapatan:")
		for _, d := range sorted {
			fmt.Printf("%s | Kategori: %s | Pendapatan: Rp%.0f\n", d.NamaSumber, d.Kategori, d.Pendapatan)
		}
	case 2:
		sorted := InsertionSortByKategori(data)
		fmt.Println("Diurutkan berdasarkan kategori:")
		for _, d := range sorted {
			fmt.Printf("%s | Kategori: %s | Pendapatan: Rp%.0f\n", d.NamaSumber, d.Kategori, d.Pendapatan)
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Menampilkan laporan bulanan penghasilan
func laporanBulanan() {
	if len(data) == 0 {
		fmt.Println("Tidak ada data untuk laporan.")
		return
	}

	var totalPendapatan, totalBiaya float64
	for _, ph := range data {
		totalPendapatan += ph.Pendapatan
		totalBiaya += ph.BiayaOperasi
	}
	laba := totalPendapatan - totalBiaya

	// Menampilkan detail per sumber
	fmt.Println("=== LAPORAN BULANAN ===")
	for i, ph := range data {
		fmt.Printf("[%d] %s | Kategori: %s | Pendapatan: Rp%.0f | Laba: Rp%.0f\n", i+1, ph.NamaSumber, ph.Kategori, ph.Pendapatan, ph.Pendapatan-ph.BiayaOperasi)
	}

	// Menampilkan total keseluruhan
	fmt.Printf("Total Pendapatan : Rp%.0f\n", totalPendapatan)
	fmt.Printf("Total Laba       : Rp%.0f\n", laba)

	// Memberi rekomendasi berdasarkan margin laba
	fmt.Println("\nRekomendasi:")
	for _, ph := range data {
		labaSumber := ph.Pendapatan - ph.BiayaOperasi
		if labaSumber > 0.5*ph.Pendapatan {
			fmt.Printf("- Pertahankan %s (Laba tinggi)\n", ph.NamaSumber)
		} else if labaSumber < 0.2*ph.Pendapatan {
			fmt.Printf("- Evaluasi %s (Laba rendah)\n", ph.NamaSumber)
		}
	}
}
