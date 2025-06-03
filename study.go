package main

import "fmt"

const NMAX int = 1024

type deadline struct {
	tahun   int
	bulan   int
	tanggal int
}

type datuma struct {
	proyek   string
	klien    string
	status   string
	bayaran  int
	deadline deadline
}

type arrData [NMAX]datuma 

func main() {
	var isiDatuma arrData
	var n, countDatuma int

	tesData(&isiDatuma, &countDatuma) 

	for n < 6 {
		menu()
		fmt.Scan(&n)
		switch n {
		case 1:
			dataLine(&isiDatuma, &countDatuma)
		case 2:
			editOption(&isiDatuma, &countDatuma)
		case 3:
			searchOption(&isiDatuma, &countDatuma)
		case 4:
			sortOption(&isiDatuma, countDatuma)
		case 5:
			showTable(&isiDatuma, countDatuma)
		case 6:
			fmt.Println("[PROGRAM SELESAI]")
		default:
			fmt.Println("[INVALID]")
		}
	}
}

func dataLine(dT *arrData, cDT *int) {
// I.S. Data array dan countData terdefinisi
// F.S. User memilih opsi menambahkan atau menghapus, fungsi add/del dijalankan
	var opsi int
	fmt.Println("\n\t\t     MENAMBAHKAN ATAU MENGHAPUS ? ")
	fmt.Println("\t\t     1. MENAMBAHKAN")
	fmt.Println("\t\t     2. MENGHAPUS")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		add(dT, cDT)
	case 2:
		del(dT, cDT)
	default:
		fmt.Println("\t\t     [INVALID]")
	}
}

func add(T *arrData, n *int) {
// I.S. Array T dan nilai n terdefinisi
// F.S. Data baru ditambahkan ke indeks ke-n dan nilai n bertambah satu
	if *n >= NMAX {
		fmt.Println("[!] Kapasitas data penuh. Tidak bisa menambah lagi.")
		return
	}// ini hanya antisipasi jika datanya penuh tapi kemungkinan tidak karena memiliki Kapasitas array yang besar

	fmt.Println("\n[INPUT DATA BARU]")
	fmt.Print("Masukkan nama proyek: ")
	fmt.Scan(&T[*n].proyek)

	fmt.Print("Masukkan nama klien: ")
	fmt.Scan(&T[*n].klien)

	fmt.Println("Masukkan deadline:")
	fmt.Print("Tahun  : ")
	fmt.Scan(&T[*n].deadline.tahun)
	fmt.Print("Bulan  : ")
	fmt.Scan(&T[*n].deadline.bulan)
	fmt.Print("Tanggal: ")
	fmt.Scan(&T[*n].deadline.tanggal)

	fmt.Print("Masukkan status: ")
	fmt.Scan(&T[*n].status)

	fmt.Print("Masukkan bayaran: ")
	fmt.Scan(&T[*n].bayaran)

	*n++
	fmt.Println("\n[✓] Data berhasil ditambahkan.")
}

func del(T *arrData, n *int) {
// I.S. Array T terdefinisi, index valid dimasukkan oleh user
// F.S. Data pada index dihapus, data setelahnya digeser, dan n berkurang satu	
	var index int
	fmt.Print("\nMasukkan indexs data yang ingin dihapus: ")
	fmt.Scan(&index)

	if index < 0 || index-1 >= *n {
		fmt.Println("[!] indexs tidak valid.")
		return
	}

	for i := index-1; i < *n-1; i++ {
		T[i] = T[i+1]
	}

	T[*n-1] = datuma{} // hapus isi yang terakhir
	*n--

	fmt.Println("\n[✓] Data berhasil dihapus.")
}

func editOption(dT *arrData, cDT *int) {
// I.S. User memilih line dan opsi edit
// F.S. Fungsi editSpecific atau editLine dijalankan sesuai pilihan user
	var opsi int
	var tanyaLine int
	fmt.Println("\n\t\t     PILIH LINE UNTUK DIEDIT: ")
	fmt.Scan(&tanyaLine)

	if tanyaLine < 0 || tanyaLine >= *cDT {
		fmt.Println("[!] Line tidak valid.")
		return
	}//jaga jaga aja kalo nguawur

	fmt.Println("\n\t\t     PILIH JENIS TIPE DATA: ")
	fmt.Println("\t\t     1. Edit Spesifik")
	fmt.Println("\t\t     2. Satu line")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		editSpecific(dT, cDT, tanyaLine-1)
	case 2:
		editLine(dT, cDT, tanyaLine-1)
	default:
		fmt.Println("\t\t     [INVALID]")
	}
}

func editLine(T *arrData, n *int, index int) {
// I.S. Array T terdefinisi, index baris valid dimasukkan
// F.S. Data pada baris ke-index diganti secara keseluruhan oleh user	
	fmt.Println("\n[EDIT SELURUH LINE]")

	fmt.Print("Masukkan nama proyek: ")
	fmt.Scan(&T[index].proyek)

	fmt.Print("Masukkan nama klien: ")
	fmt.Scan(&T[index].klien)

	fmt.Println("Masukkan deadline:")
	fmt.Print("Tahun  : ")
	fmt.Scan(&T[index].deadline.tahun)
	fmt.Print("Bulan  : ")
	fmt.Scan(&T[index].deadline.bulan)
	fmt.Print("Tanggal: ")
	fmt.Scan(&T[index].deadline.tanggal)

	fmt.Print("Masukkan status: ")
	fmt.Scan(&T[index].status)

	fmt.Print("Masukkan bayaran: ")
	fmt.Scan(&T[index].bayaran)
}

func editSpecific(T *arrData, n *int, index int) {
// I.S. Array T terdefinisi, index baris valid dan opsi kolom dipilih user
// F.S. Kolom tertentu pada baris ke-index diperbarui oleh input user	
	var opsi int
	fmt.Println("\n[EDIT SPESIFIK FIELD]")
	fmt.Println("1. Nama Proyek")
	fmt.Println("2. Nama Klien")
	fmt.Println("3. Status")
	fmt.Println("4. Bayaran")
	fmt.Println("5. Deadline")
	fmt.Print("Pilih yang ingin diedit: ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		fmt.Print("Masukkan nama proyek baru: ")
		fmt.Scan(&T[index].proyek)
	case 2:
		fmt.Print("Masukkan nama klien baru: ")
		fmt.Scan(&T[index].klien)
	case 3:
		fmt.Print("Masukkan status baru: ")
		fmt.Scan(&T[index].status)
	case 4:
		fmt.Print("Masukkan bayaran baru: ")
		fmt.Scan(&T[index].bayaran)
	case 5:
		fmt.Println("Masukkan deadline baru:")
		fmt.Print("Tahun  : ")
		fmt.Scan(&T[index].deadline.tahun)
		fmt.Print("Bulan  : ")
		fmt.Scan(&T[index].deadline.bulan)
		fmt.Print("Tanggal: ")
		fmt.Scan(&T[index].deadline.tanggal)
	default:
		fmt.Println("[!] Pilihan tidak valid.")
	}
}

func menu() {
// I.S. User memilih menu 1–6
// F.S. Fungsi sesuai pilihan dijalankan
	fmt.Println("1. EDIT LINE")
	fmt.Println("2. EDIT DATA")
	fmt.Println("3. CARI DATA")
	fmt.Println("4. URUTKAN DATA")
	fmt.Println("5. TAMPILKAN DALAM TABEL")
	fmt.Println("6. EXIT")
}

func tesData(info *arrData, tInfo *int) {
// I.S. Belum ada data contoh dalam array
// F.S. Array info berisi 5 data proyek contoh, *tInfo bernilai 5
	info[0].proyek = "Desainlogobrand"
	info[0].klien = "Amamiyaren"
	info[0].deadline = deadline{2025,6,1}
	info[0].status = "Dikerjakan"
	info[0].bayaran = 1500000
	info[1].proyek = "Websiteportofolio"
	info[1].klien = "Indrashalala"
	info[1].deadline = deadline{2025,5,20}
	info[1].status = "Selesai"
	info[1].bayaran = 800000
	info[2].proyek = "Postinganthreads"
	info[2].klien = "Sayarajen"
	info[2].deadline = deadline{2025,5,25}
	info[2].status = "Ditahan"
	info[2].bayaran = 202000
	info[3].proyek = "Schoolproject"
	info[3].klien = "User123"
	info[3].deadline = deadline{2025,6,16}
	info[3].status = "Dikerjakan"
	info[3].bayaran = 500000
	info[4].proyek = "Perawatanakun"
	info[4].klien = "Ninjaorganik"
	info[4].deadline = deadline{2025,2,8}
	info[4].status = "Selesai"
	info[4].bayaran = 489000
	*tInfo = 5
}

func searchOption(dT *arrData, cDT *int) {
// I.S. Data array dan keyword pencarian terdefinisi
// F.S. Memanggil pencarian nama secara binari atau sequential berdasarkan pilihan user	
	var opsi int
	var keyword string

	fmt.Println("\n\t\t Masukkan nama yang dicari")
	fmt.Scan(&keyword)
	fmt.Println("\n\t\t     PILIH JENIS PENCARIAN")
	fmt.Println("\t\t     1. Berdasarkan Klien")
	fmt.Println("\t\t     2. Berdasarkan Proyek")
	fmt.Print("\t\t     pilih opsi: ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		searchNamaBin(dT, cDT, keyword, "Klien")
	case 2:
		searchNamaBin(dT, cDT, keyword, "Proyek")
	default:
		fmt.Println("\t\t     [INVALID]")
	}
}

func searchNamaSeq(T *arrData, n *int, kW, by string) {
// I.S. Data array T, jumlah data n, keyword dan jenis pencarian (by) sudah ditentukan
// F.S. Menampilkan data yang sesuai keyword (klien/proyek) jika ditemukan	
	var found bool = false

	for i := 0; i < *n; i++ {
		if kW == T[i].klien || kW == T[i].proyek {
			fmt.Println("\n[DATA DITEMUKAN]")
			fmt.Printf("Proyek   : %s\n", T[i].proyek)
			fmt.Printf("Klien    : %s\n", T[i].klien)
			fmt.Printf("Deadline : %d-%d-%d\n", T[i].deadline.tahun, T[i].deadline.bulan, T[i].deadline.tanggal)
			fmt.Printf("Status   : %s\n", T[i].status)
			fmt.Printf("Bayaran  : Rp %d\n", T[i].bayaran)
			found = true
		}
	}

	if found == false {
		fmt.Println("\n[!] Data tidak ditemukan.")
	}
}

func searchNamaBin(T *arrData, n *int, kW, by string) {
// I.S. Data array proyek (T), jumlah data (n), keyword, dan jenis pencarian by "klien" atau "proyek" sudah ditentukan
// F.S. Menampilkan data proyek yang sesuai dengan keyword berdasarkan jenis pencarian jika tidak ada, tampilkan pesan tidak ditemukan

	var urut arrData
	var i, j, left, right, mid int
	var current string
	var found bool = false
	var temp datuma
	
	
	for i = 0; i < *n; i++ {// salin array ke array baru
		urut[i] = T[i]
	}
	for i = 1; i < *n; i++ {// ngurutin berdasar by pake insertion
		temp = urut[i]
		j = i - 1
		for j >= 0 && (
			(by == "Klien" && urut[j].klien > temp.klien) ||
			(by == "Proyek" && urut[j].proyek > temp.proyek)) {
			urut[j+1] = urut[j]
			j = j - 1
		}
		urut[j+1] = temp
	}
	left = 0
	right = *n - 1

	for left <= right && !found {
		mid = (left + right) / 2
		if by == "Klien" {
			current = urut[mid].klien
		} else {
			current = urut[mid].proyek
		}

		if kW == current {
			fmt.Println("\n[DATA DITEMUKAN]")
			fmt.Println("Proyek   :", urut[mid].proyek)
			fmt.Println("Klien    :", urut[mid].klien)
			fmt.Println("Deadline :", urut[mid].deadline.tahun, "-", urut[mid].deadline.bulan, "-", urut[mid].deadline.tanggal)
			fmt.Println("Status   :", urut[mid].status)
			fmt.Println("Bayaran  : Rp", urut[mid].bayaran)
			found = true
		} else {
			if kW < current {
				right = mid - 1
			} else {
				left = left + 1
			}
		}
	}

	if found == false {
		fmt.Println("\n[!] Data tidak ditemukan.")
	}
}



func showTable(dT *arrData, cDT int) {
// I.S. Data array T dan cDT terdefinisi
// F.S. Tabel proyek ditampilkan di terminal
	fmt.Println("\n===================================== TABEL DATA PROYEK =====================================")
	fmt.Printf("%-4s %-25s %-20s %-12s %-20s %-10s\n", "No", "Proyek", "Klien", "Deadline", "Status", "Bayaran")
	fmt.Println("---------------------------------------------------------------------------------------------")

	for i := 0; i < cDT; i++ {
		fmt.Printf("%-4d %-25s %-20s %04d-%02d-%02d   %-20s Rp %-10d\n",
			i+1,dT[i].proyek,dT[i].klien,dT[i].deadline.tahun,dT[i].deadline.bulan,dT[i].deadline.tanggal,dT[i].status,dT[i].bayaran)
	}
	fmt.Println("=============================================================================================")
}


func sortOption(dT *arrData, cDT int) {
// I.S. User memilih opsi sorting, array T dan countData sudah terisi
// F.S. Data terurut dan tabel ditampilkan	
	var opsi int
	fmt.Println("\n\t\t     OPSI SORT DATA")
	fmt.Println("\t\t     1. Berdasarkan Deadline[kecil -> terbesar]")
	fmt.Println("\t\t     2. Berdasarkan Deadline[besar -> terkecil]")
	fmt.Println("\t\t     3. Berdasarkan Bayaran[kecil -> terbesar]")
	fmt.Println("\t\t     4. Berdasarkan Bayaran[besar -> terkecil]")
	fmt.Print("\t\t     pilih opsi: ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		typeSortInsertion(dT, &cDT, "deadline_asc")
	case 2:
		typeSortInsertion(dT, &cDT, "deadline_desc")
	case 3:
		typeSortSelection(dT, &cDT, "bayaran_asc")
	case 4:
		typeSortInsertion(dT, &cDT, "bayaran_desc")
	default:
		fmt.Println("\t\t     [INVALID]")
	}
}

func convertDate(date deadline) int {
	// mengubah dari 2025-10-21 jadi 20251021
	return date.tahun*10000 + date.bulan*100 + date.tanggal
}

func typeSortInsertion(T *arrData, n *int, by string) {
// I.S. Array T dan jumlah data n terdefinisi
// F.S. Array T terurut sesuai mode (deadline/bayaran naik/turun)	
	var temp datuma
	var i, j int
	var swap bool

	for i = 1; i < *n; i++ {
		temp = T[i]
		j = i - 1
		swap = true

		for j >= 0 && swap {
			swap = false

			if by == "deadline_asc" && convertDate(T[j].deadline) > convertDate(temp.deadline) {
				swap = true
			} else if by == "deadline_desc" && convertDate(T[j].deadline) < convertDate(temp.deadline) {
				swap = true
			} else if by == "bayaran_asc" && T[j].bayaran > temp.bayaran {
				swap = true
			} else if by == "bayaran_desc" && T[j].bayaran < temp.bayaran {
				swap = true
			}

			if swap {
				T[j+1] = T[j]
				j = j - 1
			}
		}
		T[j+1] = temp
	}

	fmt.Println("\nData berhasil diurutkan berdasarkan", by)
	showTable(T, *n)
}



func typeSortSelection(T *arrData, n *int, by string) {
// I.S. Array T dan jumlah data n terdefinisi
// F.S. Array T terurut sesuai mode (deadline/bayaran naik/turun)	
	var i, j, idx int
	var temp datuma
	var swap bool

	for i = 0; i < *n; i++ {
		idx = i
		for j = i + 1; j < *n; j++ {
			swap = false

			if by == "deadline_asc" && convertDate(T[j].deadline) < convertDate(T[idx].deadline) {
				swap = true
			} else if by == "deadline_desc" && convertDate(T[j].deadline) > convertDate(T[idx].deadline) {
				swap = true
			} else if by == "bayaran_asc" && T[j].bayaran < T[idx].bayaran {
				swap = true
			} else if by == "bayaran_desc" && T[j].bayaran > T[idx].bayaran {
				swap = true
			}

			if swap {
				idx = j
			}
		}
		if idx != i {
			temp = T[i]
			T[i] = T[idx]
			T[idx] = temp
		}
	}

	fmt.Println("\nData berhasil diurutkan berdasarkan", by)
	showTable(T, *n)
}
