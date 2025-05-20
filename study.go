package main

import"fmt"

const NMAX int = 1024

type deadline struct{
	tahun    int
	bulan    int
	tanggal  int
}

type datuma struct{
	proyek   string
	klien    string
	status   string
	bayaran  int
	deadline deadline 
}

type arrData[NMAX-1]datuma

func main(){
	var isiDatuma arrData
	var n, countDatuma int
	
	tesData(isiDatuma, &countDatuma)
	
	for n < 6 {
		menu()
		fmt.Scan(&n)
		switch n {
			case 1:
				tambahData(...)
			case 2:
				editOption(isiDatuma, &countDatuma)
			case 3:
				searchOption(&isiDatuma, &countDatuma)//udh
			case 4:
				sortOption(&isiDatuma, &countDatuma)//spek terpenuhi, pelajari lagi
			case 5:
				showTable(&isiDatuma, countDatuma)//sesai
			case 6:
				fmt.Println("[PROGRAM SELESAI]")
			default:
				fmt.Println("[INVALID]")
		}
	}
}

func editOption(dT *arrData, cDT *int){
	var opsi int
	var tanyaLine arrData
	fmt.Println("\n\t\t     PILIH LINE UNTUK DIEDIT: ")
	fmt.Scan(&tanyaLine[NMAX-1])
	fmt.Println("\n\t\t     PILIH JENIS TIPE DATA: ")
	fmt.Println("\t\t     1. Edit Spesifik")
	fmt.Println("\t\t     2. Satu line")
	fmt.Scan(&opsi)
	
	switch opsi{
		case 1:
			editSpecific(dT,cDT,"Proyek")
		case 2:
			editLine(...)
		default:
			fmt.Println("\t\t     [INVALID]")
	}
}

func editLine

func editSpecific(T *arrData, n *int){
	
	
}



func menu(){
	fmt.Println("1. TAMBAHKAN DATA")//tambah line baru di data
	fmt.Println("2. EDIT DATA")//isinya PERBAHARUI DATA(update status) dan HAPUS DATA(bisa satu line di hapus dari tabel)
	fmt.Println("3. CARI DATA")//cari data berdasarkan klien atau proyek
	fmt.Println("4. URUTKAN DATA")//urutkan berdasarkan deadline atau bayaran dari tertinggi (ya dari terendah sekalian)
	fmt.Println("5. TAMPILKAN DALAM TABEL")//ya tampilin tabel biar enak diliat
	fmt.Println("6. EXIT")
}

func tesData(info *arrData, tInfo *int){
	info[0].proyek = "Desain Logo Brand"
	info[0].klien = "Amamiya Ren"
	info[0].deadline = deadline{2025,6,1}
	info[0].status = "Sedang dikerjakan"
	info[0].bayaran = 1500000
	info[1].proyek = "Website Portfolio"
	info[1].klien = "Indra Shalala"
	info[1].deadline = deadline{2025,5,20}
	info[1].status = "Selesai"
	info[1].bayaran = 800000
	info[2].proyek = "Postingan Threads"
	info[2].klien = "sayaRajen"
	info[2].deadline = deadline{2025,5,25}
	info[2].status = "Pending"
	info[2].bayaran = 202000
	info[3].proyek = "School Project"
	info[3].klien = "user123"
	info[3].deadline = deadline{2025,6,16}
	info[3].status = "Sedang dikerjakan"
	info[3].bayaran = 500000
	info[4].proyek = "perawatan akun"
	info[4].klien = "NinjaOrganik"
	info[4].deadline = "2025-02-08"
	info[4].deadline = deadline{2025,2,8}
	info[4].status = "selesai"
	info[4].bayaran = 489000
	*tInfo = 5
}

func searchOption(dT *arrData, cDT *int){
	var opsi int
	var keyword string
	
	fmt.Println("\n\t\t Masukkan nama yang dicari")
	fmt.Scan(&keyword)
	fmt.Println("\n\t\t     PILIH JENIS PENCARIAN")
	fmt.Println("\t\t     1.Berdasarkan Klien")
	fmt.Println("\t\t     2.Berdasarkan Proyek")
	fmt.Print("\t\t     pilih opsi: ")
	fmt.Scan(&opsi)
	
	switch opsi {
		case 1:
			searchNamaBin(dT, cDT, keyword, "Klien")
		case 2:
			searchNamaSeq(dT, cDT, keyword, "Proyek")
		default:
			fmt.Println("\t\t     [INVALID]")
	}
}

func searchNamaSeq ( T *arrData, n *int,kW, by string){
	var found bool = false
	var temp string
	var i int
	var j int = 0
	
	for i = 0; i < n; i++{
		if kW == T[i].klien{
			//var j int = 0
			for j < n && !found{
				found = T[i].klien
				if found==false{
					fmt.Println("\n[!] Data tidak ditemukan.")
				}else if found==true{
		
					fmt.Println("\n[DATA DITEMUKAN]")
					fmt.Printf("Proyek   : %s\n", T[i].proyek)
					fmt.Printf("Klien    : %s\n", T[i].klien)
					fmt.Printf("Deadline : %d-%d-%d\n", T[i].deadline.tahun,T[i].deadline.bulan,T[i].deadline.tanggal)
					fmt.Printf("Status   : %s\n", T[i].status)
					fmt.Printf("Bayaran  : Rp %d\n", T[i].bayaran)
					found = true
					
				}
			}
		}else if kW == T[i].proyek{
			//var j int = 0
			for j < n && !found{
				found = T[i].proyek
				if found==false{
					fmt.Println("\n[!] Data tidak ditemukan.")
				}else { //tidak pakai if found==true
		
					fmt.Println("\n[DATA DITEMUKAN]")
					fmt.Printf("Proyek   : %s\n", T[i].proyek)
					fmt.Printf("Klien    : %s\n", T[i].klien)
					fmt.Printf("Deadline : %d-%d-%d\n", T[i].deadline.tahun,T[i].deadline.bulan,T[i].deadline.tanggal)
					fmt.Printf("Status   : %s\n", T[i].status)
					fmt.Printf("Bayaran  : Rp %d\n", T[i].bayaran)
					found = true
					
				}
			}
		}
	}
}

func searchNamaBin(T *arrData, n *int, kW, by string) {
	var left, right, mid int
	var current string
	var found bool = false

	left = 0
	right = *n - 1

	for left <= right && !found {
		mid = (left + right) / 2

		if by == "Klien" {
			current = T[mid].klien
		} else if by == "Proyek" {
			current = T[mid].proyek
		} else {
			fmt.Println("[!] Jenis pencarian tidak valid.")
			return
		}

		if keyword == current {
			fmt.Println("\n[DATA DITEMUKAN]")
			fmt.Println("Proyek   :", T[mid].proyek)
			fmt.Println("Klien    :", T[mid].klien)
			fmt.Println("Deadline :", T[mid].deadline.tahun, "-", T[mid].deadline.bulan, "-", T[mid].deadline.tanggal)
			fmt.Println("Status   :", T[mid].status)
			fmt.Println("Bayaran  : Rp", T[mid].bayaran)
			found = true
		} else if keyword < current {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if !found {
		fmt.Println("\n[!] Data tidak ditemukan.")
	}
}




func showTable(dT *arrData, cDT int){
	var i int 
	var found bool
	
	fmt.Println("\n==================== TABEL DATA PROYEK ====================")
	fmt.Println("No\tProyek\tKlien\tDeadline\tStatus\tBayaran")
	fmt.Println("--------------------------------------------------------------------")

	for i = 0; i < NMAX; i++{
		if dT[i].proyek == ""{
			found = true
		fmt.Println(i+1,"\t",dT[i].proyek,"\t",dT[i].klien,"\t",dT[i].deadline,"\t",dT[i].status,"\t",dT[i].bayaran)
		}
		
	}
}

func sortOption(dT *arrData, cDT int){
	var opsi int
	fmt.Println("\n\t\t     OPSI SORT DATA")
	fmt.Println("\t\t     1. Berdasarkan Deadline[kecil -> terbesar]")
	fmt.Println("\t\t     2. Berdasarkan Deadline[besar -> terkecil]")
	fmt.Println("\t\t     3. Berdasarkan Bayaran[kecil -> terbesar]")
	fmt.Println("\t\t     4. Berdasarkan Bayaran[besar -> terkecil]")
	fmt.Print("\t\t     pilih opsi: ")
	fmt.Scan(&opsi)
	
	switch opsi{
		case 1:
			typeSortInsertion(dT, *cDT, "deadline_asc")
		case 2:
			typeSortInsertion(dT, *cDT, "deadline_desc")
		case 3:
			typeSortSelection(dT, *cDT, "bayaran_asc")
		case 4:
			typeSortInsertion(dT, *cDT, "bayaran_desc")
		default:
			fmt.Println("\t\t     [INVALID]")
	}
}

func convertDate(date deadline)int{
	//mengubah dari yang seperti 2025-05-14 menjadi 20250514
	return d.tahun*10000+d.bulan*100+d.tanggal
}

func typeSortInsertion(T *arrData, n *int, by string){
	var i, j int
	var temp datuma
	var swap bool 
	
	for i = 1;i < n;i++{
		temp = T[i]
		j = i - 1
	
		for j >= 0 {
			swap = false
				
			if by == "deadline_asc" {
				if convertDate(T[j].deadline) > convertDate(T[j+1].deadline) {
					swap = true
				}
			} else if by == "deadline_desc" {
				if convertDate(T[j].deadline) < convertDate(T[j+1].deadline) {
					swap = true
				}
			} else if by == "bayaran_asc" {
				if T[j].bayaran > T[j+1].bayaran {
					swap = true
				}
			} else if by == "bayaran_desc" {
				if T[j].bayaran < T[j+1].bayaran {
					swap = true
				}
			}

			if swap {
				T[j+1] = T[j]
				j--
			}else {
				j = -1
			}
		}
		T[j+1]=temp
	}
	fmt.Println("\n Data berhasil diurutkan berdasarkan",by)
	showTable(T, *n)
}

func typeSortSelection(T *arrData, n *int, b string){
	var i, min, j int
	var temp datuma
	for i = 0; i < *n;i++{
		min = i
		for j = i+1; j < *n;j++{
			///////
			if by == "deadline_asc"{
				if convertDate(T[j].deadline) < convertDate(T[min].deadline){
					min = j
				}
			}else if by == "deadline_desc"{
				if convertDate(T[j].deadline) < convertDate(T[min].deadline){
					min = j
				}
			}else if by == "bayaran_asc"{
				if T[j].bayaran < T[min].bayaran{
					min = j
				}
			}else if by == "bayaran_desc"{
				if T[j].bayaran < T[min].bayaran{
					min = j
				}
			}
		}
		if min != i{
			temp = T[min]
			T[min] = T[i]
			T[i] = temp
		}
	}
	fmt.Println("\n Data berhasil diurutkan berdasarkan",by)
	showTable(T, *n)
}
