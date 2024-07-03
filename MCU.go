package main

import "fmt"

const NMAX int = 100

type pasien struct {
	nama  string
	tahun int
	jenis string
	harga int
	jantung, gulDar, urine, ginjal, kolestrol string
}

type tabPasien [NMAX]pasien

var count int

func main() {
	var find bool
	var x, n, y, z, a int
	var data tabPasien
	var choose string
	var chooseAD string
	find = false
	x = menuUtama()

	for x != 1 && x != 4 {
		fmt.Println()
		fmt.Println("Masukan data terlebih dahulu")
		fmt.Println()
		x = menuUtama()
	}
	for x != 4 {
		if x == 1 {
			fmt.Println()
			fmt.Println(">>>>>>>>>> MASUKAN DATA <<<<<<<<<<")
			fmt.Println()
			if find{
				fmt.Println("Anda sudah memasukkan data awal. Silahkan tambahkan di menu tambah data!")
				fmt.Println()
		}else {
			fmt.Print("Masukkan banyak data: ")
			fmt.Scan(&n)
			fmt.Println()
			masukanData(&data, &n, &count)
			find = true
		}
		} else if x == 2 {

			fmt.Println()
			fmt.Println(">>>>>>>>>> TAMBAH, EDIIT, HAPUS <<<<<<<<<<")
			fmt.Println()
			y = TEHdata()

			if y == 1 {
				fmt.Println()
				fmt.Println(">>>>>>>>>> TAMBAH DATA <<<<<<<<<<")
				fmt.Println()
				tambahData(&data, &count)
			} else if y == 2 {
				fmt.Println()
				fmt.Println(">>>>>>>>>> EDIIT DATA <<<<<<<<<<")
				fmt.Println()
				editData(&data, &count)
			} else if y == 3 {
				fmt.Println()
				fmt.Println(">>>>>>>>>> HAPUS DATA <<<<<<<<<<")
				fmt.Println()
				hapusData(&data, &count)
			} else if y == 4 {
				menuUtama()
			} else if y == 5 {
				fmt.Print("Terima Kasih!")
			}
		} else if x == 3 {
			fmt.Println()
			fmt.Println(">>>>>>>>>> INFO DATA <<<<<<<<<<")
			fmt.Println()
			z = infoData()
			if z == 1 {
				fmt.Println()
				fmt.Println(">>>>>>>>>> INFO PEMASUKAN <<<<<<<<<<")
				fmt.Println()
				infoPemasukan(&data, &count)
			} else if z == 2 {
				fmt.Println()
				fmt.Println(">>>>>>>>>> CARI PASIEN <<<<<<<<<<")
				fmt.Println()
				a = CariPasien()
				if a == 1 {
					filterPaket(&data, &count)
				} else if a == 2 {
					filterTahun(&data, &count)
				} else if a == 3 {
					filterNama(&data, &count)
				}
			} else if z == 3 {
				fmt.Println()
				fmt.Println(">>>>>>>>>> URUTKAN PASIEN <<<<<<<<<<")
				fmt.Println()
				fmt.Print("Pilih berdasarkan(Tahun / Paket): ")
				fmt.Scan(&choose)
				fmt.Print("Urut Menaik / Menurun: ")
				fmt.Scan(&chooseAD)
				if choose == "Tahun" && chooseAD == "Menaik" {
					urutkanPasienByTahunAscending(&data, &count)
				} else if choose == "Tahun" && chooseAD == "Menurun" {
					urutkanPasienByTahunDescending(&data, &count)
				} else if choose == "Paket" && chooseAD == "Menaik" {
					urutkanPasienByPaketAscending(&data, &count)
				} else if choose == "Paket" && chooseAD == "Menurun" {
					urutkanPasienByPaketDescending(&data, &count)
				}
			} else if z == 4{
				fmt.Println()
				fmt.Println(">>>>>>>>>> REKAP DATA PASIEN <<<<<<<<<<")
				fmt.Println()
				rekap(&data, &count)
			}
		} else if x == 4 {
			fmt.Println("Terima Kasih!")
		}
		x = menuUtama()
	}
}

func menuUtama() int {
	var pilih int

	fmt.Println("||====================================||")
	fmt.Println("||      APLIKASI MEDICAL CHECKUP      ||")
	fmt.Println("||====================================||")
	fmt.Println("||              >> Menu               ||")
	fmt.Println("||------------------------------------||")
	fmt.Println("||  1.   Masukan Data                 ||")
	fmt.Println("||  2.   Tambah, Edit dan Hapus Data  ||")
	fmt.Println("||  3.   Info Data                    ||")
	fmt.Println("||  4.   Exit                         ||")
	fmt.Println("||------------------------------------||")
	fmt.Print("Pilih opsi (1/2/3/4): ")
	fmt.Scan(&pilih)
	return pilih

}

func masukanData(A *tabPasien, n *int, count *int) {

	for i := 0; i < *n; i++ {
		fmt.Println("Masukan data (Nama, paket, harga dan tahun):")
		fmt.Println()
		fmt.Print("Nama Pasien (ex. Amir): ")
		fmt.Scan(&A[i].nama)
		fmt.Print("Paket Pasien (Silver/Gold/Platinum): ")
		fmt.Scan(&A[i].jenis)
		for A[i].jenis != "Silver" && A[i].jenis != "Gold" && A[i].jenis != "Platinum"{
			fmt.Println()
			fmt.Println("Masukan salah! Paket : Silver/Gold/Platinum !")
			fmt.Print("Masukan paket (Silver/Gold/Platinum): ")
			fmt.Scan(&A[i].jenis)
		}
		fmt.Print("Harga Paket (Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
		fmt.Scan(&A[i].harga)
		for A[i].jenis == "Silver" && A[i].harga != 100000{
			fmt.Println()
			fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
			fmt.Print("(Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
			fmt.Scan(&A[i].harga)
		}
		for A[i].jenis == "Gold" && A[i].harga != 250000{
			fmt.Println()
			fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
			fmt.Print("(Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
			fmt.Scan(&A[i].harga)
		}
		for A[i].jenis == "Platinum" && A[i].harga != 400000{
			fmt.Println()
			fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
			fmt.Print("(Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
			fmt.Scan(&A[i].harga)
		}
		fmt.Print("Periode Tahun (ex. 2023): ")
		fmt.Scan(&A[i].tahun)
		fmt.Println()

		fmt.Print("           Form Kondisi Pasien           \n")

		fmt.Print("Kondisi jantung : ")
		fmt.Scan(&A[i].jantung)
		for A[i].jantung != "Normal" && A[i].jantung != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!\n")

			fmt.Print("Kondisi jantung : ")
			fmt.Scan(&A[i].jantung)
		}

		fmt.Print("Kondisi gula darah : ")
		fmt.Scan(&A[i].gulDar)
		for A[i].gulDar != "Normal" && A[i].gulDar != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!\n")

			fmt.Print("Kondisi gula darah : ")
			fmt.Scan(&A[i].gulDar)
		}

		fmt.Print("Kondisi urine : ")
		fmt.Scan(&A[i].urine)
		for A[i].urine != "Normal" && A[i].urine != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!\n")

			fmt.Print("Kondisi urine : ")
			fmt.Scan(&A[i].urine)
		}

		fmt.Print("Kondisi ginjal : ")
		fmt.Scan(&A[i].ginjal)
		for A[i].ginjal != "Normal" && A[i].ginjal != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!\n")

			fmt.Print("Kondisi ginjal : ")
			fmt.Scan(&A[i].ginjal)
		}

		fmt.Print("Kondisi kolestrol : ")
		fmt.Scan(&A[i].kolestrol)
		for A[i].kolestrol != "Normal" && A[i].kolestrol != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!\n")

			fmt.Print("Kondisi kolestrol : ")
			fmt.Scan(&A[i].kolestrol)
		}
		*count++
	}

	fmt.Println()
	fmt.Print("Data :\n")
	for i := 0; i < *count; i++ {
		fmt.Print(i+1, ". ")
		fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
	}
	fmt.Println()
}

func TEHdata() int {
	var pilih int

	fmt.Println("||====================================||")
	fmt.Println("||      APLIKASI MEDICAL CHECKUP      ||")
	fmt.Println("||====================================||")
	fmt.Println("||       >> Tambah, Edit, Hapus       ||")
	fmt.Println("||------------------------------------||")
	fmt.Println("||  1.   Tambah Data                  ||")
	fmt.Println("||  2.   Edit Data                    ||")
	fmt.Println("||  3.   Hapus Data                   ||")
	fmt.Println("||  4.   Menu                         ||")
	fmt.Println("||  5.   Exit                         ||")
	fmt.Println("||------------------------------------||")
	fmt.Print("Pilih opsi (1/2/3/4/5): ")
	fmt.Scan(&pilih)
	return pilih
}

func tambahData(A *tabPasien, count *int) {
	var tambah int
	var temp, i int
	fmt.Print("Jumlah data yang ingin ditambah :")
	fmt.Scan(&tambah)
	fmt.Println()
	temp = *count
	i = *count
	fmt.Print("Masukkan data yang akan ditambah")
	fmt.Println()
	fmt.Print("(Nama, Paket, Harga, Tahun,) :\n")
	for i < tambah+temp {
		fmt.Print("Nama Pasien (ex. Amir): ")
		fmt.Scan(&A[i].nama)
		fmt.Print("Paket Pasien (Silver/Gold/Platinum): ")
		fmt.Scan(&A[i].jenis)
		for A[i].jenis != "Silver" && A[i].jenis != "Gold" && A[i].jenis != "Platinum"{
			fmt.Println()
			fmt.Println("Masukan salah! Paket : Silver/Gold/Platinum !")
			fmt.Print("Masukan paket (Silver/Gold/Platinum): ")
			fmt.Scan(&A[i].jenis)
		}
		fmt.Print("Harga Paket (Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
		fmt.Scan(&A[i].harga)
		for A[i].jenis == "Silver" && A[i].harga != 100000{
			fmt.Println()
			fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
			fmt.Print("(Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
			fmt.Scan(&A[i].harga)
		}
		for A[i].jenis == "Gold" && A[i].harga != 250000{
			fmt.Println()
			fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
			fmt.Print("(Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
			fmt.Scan(&A[i].harga)
		}
		for A[i].jenis == "Platinum" && A[i].harga != 400000{
			fmt.Println()
			fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
			fmt.Print("(Masukan Harga sesuai paket pasien {Silver[100000]/Gold[250000]/Platinum[400000]}): ")
			fmt.Scan(&A[i].harga)
		}
		fmt.Print("Periode Tahun (ex. 2023): ")
		fmt.Scan(&A[i].tahun)

		fmt.Println()
		fmt.Print("           Form Kondisi Pasien           ")
		fmt.Println()

		fmt.Print("Kondisi Jantung : ")
		fmt.Scan(&A[i].jantung)
		for A[i].jantung != "Normal" && A[i].jantung != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!")
			fmt.Println()
			fmt.Print("Kondisi jantung : ")
			fmt.Scan(&A[i].jantung)
		}

		fmt.Print("Kondisi Gula Darah : ")
		fmt.Scan(&A[i].gulDar)
		for A[i].gulDar != "Normal" && A[i].gulDar != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!")
			fmt.Println()
			fmt.Print("Kondisi Gula Darah : ")
			fmt.Scan(&A[i].gulDar)
		}

		fmt.Print("Kondisi Urine : ")
		fmt.Scan(&A[i].urine)
		for A[i].urine != "Normal" && A[i].urine != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!")
			fmt.Println()
			fmt.Print("Kondisi Urine : ")
			fmt.Scan(&A[i].urine)
		}

		fmt.Print("Kondisi Ginjal : ")
		fmt.Scan(&A[i].ginjal)
		for A[i].ginjal != "Normal" && A[i].ginjal != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!")
			fmt.Println()
			fmt.Print("Kondisi Ginjal : ")
			fmt.Scan(&A[i].ginjal)
		}

		fmt.Print("Kondisi kolestrol : ")
		fmt.Scan(&A[i].kolestrol)
		for A[i].kolestrol != "Normal" && A[i].kolestrol != "Tidak_Normal"{
			fmt.Println()
			fmt.Print("Masukan antara Normal atau Tidak_Normal. Terima kasih!\n")

			fmt.Print("Kondisi kolestrol : ")
			fmt.Scan(&A[i].kolestrol)
		}
		i++
		*count++
	}
	fmt.Println()
	fmt.Print("Data baru: ")
	fmt.Println()
	for i := 0; i < *count; i++ {
		fmt.Print(i+1, ". ")
		fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
	}
	fmt.Println()
}

// sequential search
func editData(A *tabPasien, count *int) {
	var targetNama string
	var targetTahun int
	var newNama string
	var newTahun int
	var newJenis, newJantung, newGuldar, newUrine, newGinjal, newKolestrol string
	var newHarga int
	var edit, editKondisi string

	fmt.Print("Masukkan nama pasien yang ingin diedit (ex. Jhony): ")
	fmt.Scan(&targetNama)
	fmt.Print("Masukkan tahun pasien yang ingin diedit (ex. 2022): ")
	fmt.Scan(&targetTahun)

	found := false
	for i := 0; i < *count; i++ {
		if A[i].nama == targetNama && A[i].tahun == targetTahun {
			found = true
			fmt.Println()
			fmt.Println("Data ditemukan. Ingin mengubah apa?")
			fmt.Print("(Nama/Tahun/Paket/Kondisi/Semua): ")
			fmt.Scan(&edit)
			fmt.Println()
			if edit == "Nama" {
				fmt.Print("Nama baru: ")
				fmt.Scan(&newNama)
				A[i].nama = newNama
				fmt.Println()
				} else if edit == "Tahun" {
					fmt.Print("Tahun baru: ")
					fmt.Scan(&newTahun)
					A[i].tahun = newTahun
					fmt.Println()
				} else if edit == "Paket"{
					fmt.Print("Paket baru: ")
					fmt.Scan(&newJenis)
					fmt.Print("Harga paket: ")
					fmt.Scan(&newHarga)
					for newJenis == "Silver" && newHarga != 100000{
						fmt.Println()
						fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
						fmt.Println("(Masukan Harga dan Paket yang sesuai!): ")
						fmt.Print("Paket Baru: ")
						fmt.Scan(&newJenis)
						fmt.Print("Harga Paket: ")
						fmt.Scan(&newHarga)
					}
					for newJenis == "Gold" && newHarga != 250000{
						fmt.Println()
						fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
						fmt.Println("(Masukan Harga dan Paket yang sesuai!): ")
						fmt.Print("Paket Baru: ")
						fmt.Scan(&newJenis)
						fmt.Print("Harga Paket: ")
						fmt.Scan(&newHarga)
					}
					for newJenis == "Platinum" && newHarga != 400000{
						fmt.Println()
						fmt.Println("Harga Salah! Harga paket : Silver[100000]/Gold[250000]/Platinum[400000]!")
						fmt.Println("(Masukan Harga dan Paket yang sesuai!): ")
						fmt.Print("Paket Baru: ")
						fmt.Scan(&newJenis)
						fmt.Print("Harga Paket: ")
						fmt.Scan(&newHarga)
					}
					A[i].jenis = newJenis
					A[i].harga = newHarga
					fmt.Println()
				} else if edit == "Kondisi"{
					fmt.Println("Update kondisi (Jantung/Kondisi_Gula_Darah/Kondisi_Urine/Kondisi_Ginjal/Kondisi_Kolestrol): ")
					fmt.Scan(&editKondisi)
					fmt.Println()
					if editKondisi == "Jantung"{
						fmt.Print("Update Kondisi Jantung: ")
						fmt.Scan(&newJantung)
						A[i].jantung = newJantung
					} else if edit == "Kondisi_Gula_Darah"{
						fmt.Print("Update kondisi : ")
						fmt.Scan(&newGuldar)
						A[i].gulDar = newGuldar
					} else if edit == "Kondisi_Urine"{
						fmt.Print("Update kondisi : ")
						fmt.Scan(&newUrine)
						A[i].urine = newUrine
					} else if edit == "Kondisi_Ginjal"{
						fmt.Print("Update kondisi : ")
						fmt.Scan(&newGinjal)
						A[i].ginjal = newGinjal
					} else if edit == "Kondisi_Kolestrol"{
						fmt.Print("Update Kondisi : ")
						fmt.Scan(&newKolestrol)
						A[i].kolestrol = newKolestrol
					}
				} else if edit == "Semua" {
					fmt.Print("Nama baru: ")
					fmt.Scan(&newNama)
					fmt.Print("Tahun baru: ")
					fmt.Scan(&newTahun)
					fmt.Print("Paket baru: ")
					fmt.Scan(&newJenis)
					fmt.Print("Harga baru: ")
					fmt.Scan(&newHarga)
					fmt.Print("Update jantung: ")
					fmt.Scan(&newJantung)
					fmt.Print("Update gula darah: ")
					fmt.Scan(&newGuldar)
					fmt.Print("Update urine: ")
					fmt.Scan(&newUrine)
					fmt.Print("Update ginjal: ")
					fmt.Scan(&newGinjal)
					fmt.Print("Update kolestrol: ")
					fmt.Scan(&newKolestrol)
					A[i].nama = newNama
					A[i].tahun = newTahun
					A[i].jenis = newJenis
					A[i].harga = newHarga
					A[i].jantung = newJantung
					A[i].gulDar = newGuldar
					A[i].urine = newUrine
					A[i].ginjal = newGinjal
					A[i].kolestrol = newKolestrol
				}
			fmt.Println()
			fmt.Println("Data berhasil diubah.")
			fmt.Println()
			fmt.Println("Data setelah di edit: ")
			for i := 0; i < *count; i++ {
				fmt.Print(i+1, ". ")
				fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
			}
			fmt.Println()
		}

	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
		fmt.Println()
	}	
}

func hapusData(A *tabPasien, count *int) {
	var nama string
	found := false
	fmt.Print("Masukan nama yang akan dihapus datanya: ")
	fmt.Scan(&nama)
	for i := *count - 1; i >= 0; i-- {
		if A[i].nama == nama {
			found = true
			for j := i; j < *count-1; j++ {
				A[j] = A[j+1]
			}
			*count--
			fmt.Println()
			fmt.Printf("Data untuk %s telah dihapus.\n", nama)
			for i := 0; i < *count; i++ {
				fmt.Print(i+1, ". ")
				fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
			}
			fmt.Println()
		}
	}
	if !found {
		fmt.Printf("Data untuk %s tidak ditemukan.\n", nama)
	}
}

func infoData() int {
	var pilih int
	
	fmt.Println("||====================================||")
	fmt.Println("||      APLIKASI MEDICAL CHECKUP      ||")
	fmt.Println("||====================================||")
	fmt.Println("||            >> Info Data            ||")
	fmt.Println("||------------------------------------||")
	fmt.Println("||  1.   Info Pemasukan               ||")
	fmt.Println("||  2.   Cari Pasien                  ||")
	fmt.Println("||  3.   Urutkan Pasien               ||")
	fmt.Println("||  4.   Rekap                        ||")
	fmt.Println("||  5.   Menu                         ||")
	fmt.Println("||  6.   Exit                         ||")
	fmt.Println("||------------------------------------||")
	fmt.Print("Pilih opsi (1/2/3/4/5/6): ")
	fmt.Scan(&pilih)
	return pilih
	
}

func infoPemasukan(A *tabPasien, count *int) {
	var tahun int
	
	fmt.Print("Masukkan Periode Tahun: ")
	fmt.Scan(&tahun)
	fmt.Println()
	fmt.Println("Data pada tahun ",tahun," :")
	total := 0
	for i := 0; i < *count; i++ {
		if A[i].tahun == tahun {
			fmt.Print(i+1, ". ")
			fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
		}
	}
	fmt.Println()
	for i := *count - 1; i >= 0; i-- {
		if A[i].tahun == tahun {
			total += A[i].harga
		}
	}
	fmt.Println()
	fmt.Println("Pemasukan pada periode tahun ", tahun, " adalah: ", total)
	fmt.Println()
}

func CariPasien() int {
	var pilih int

	fmt.Println("||====================================||")
	fmt.Println("||      APLIKASI MEDICAL CHECKUP      ||")
	fmt.Println("||====================================||")
	fmt.Println("||           >> Cari Pasien           ||")
	fmt.Println("||------------------------------------||")
	fmt.Println("||  1.   Filter Paket                 ||")
	fmt.Println("||  2.   Filter Tahun                 ||")
	fmt.Println("||  3.   Filter Nama                  ||")
	fmt.Println("||  4.   Back                         ||")
	fmt.Println("||  5.   Exit                         ||")
	fmt.Println("||------------------------------------||")
	fmt.Print("Pilih opsi (1/2/3/4/5): ")
	fmt.Scan(&pilih)
	fmt.Println()
	return pilih
}
func urutkanPaket(A *tabPasien, count *int) {
	for i := 0; i < *count-1; i++ {
		for j := 0; j < *count-i-1; j++ {
			if A[j].harga > A[j+1].harga {
				temp := A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}
		}
	}
}
func urutkanTahun(A *tabPasien, count *int) {
	for i := 0; i < *count-1; i++ {
		for j := 0; j < *count-i-1; j++ {
			if A[j].tahun > A[j+1].tahun {
				temp := A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}
		}
	}
}

func urutkanNama(A *tabPasien, count *int) {
	for i := 0; i < *count-1; i++ {
		for j := 0; j < *count-i-1; j++ {
			if A[j].tahun > A[j+1].tahun {
				temp := A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}
		}
	}
}

func filterPaket(A *tabPasien, count *int) {
	var pilih string
	fmt.Print("Daftar pasien dengan paket (Silver/Gold/Platinum): ")
	fmt.Scan(&pilih)

	// urut data pasien berdasarkan paket
	urutkanPaket(A, count)

	// binary search
	low := 0
	high := *count - 1
	start := -1
	end := -1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if A[mid].jenis == pilih {
			start = mid
			for start >= 0 && A[start].jenis == pilih {
				start--
			}
			end = mid
			for end < *count && A[end].jenis == pilih {
				end++
			}
			start++
			found = true
			break
		} else if A[mid].jenis < pilih {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		for i := start; i < end; i++ {
			fmt.Printf("%d. %s %s %d %d\n", i-start+1, A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
		}
	} else {
		fmt.Printf("Tidak ada pasien dengan paket %s\n", pilih)
	}
}

func filterTahun(A *tabPasien, count *int) {
	var pilih int
	fmt.Print("Daftar pasien dengan tahun: ")
	fmt.Scan(&pilih)

	// rut data pasien berdasarkan tahun
	urutkanTahun(A, count)

	// binary search
	low := 0
	high := *count - 1
	start, end := -1, -1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if A[mid].tahun == pilih {
			start = mid
			for start >= 0 && A[start].tahun == pilih {
				start--
			}
			end = mid
			for end < *count && A[end].tahun == pilih {
				end++
			}
			start++
			found = true
			break
		} else if A[mid].tahun < pilih {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		for i := start; i < end; i++ {
			fmt.Printf("%d. %s %s %d %d\n", i-start+1, A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
		}
	} else {
		fmt.Printf("Tidak ada pasien dengan tahun %d\n", pilih)
	}
}

func filterNama(A *tabPasien, count *int) {
	var pilih string
	fmt.Print("Daftar pasien dengan nama: ")
	fmt.Scan(&pilih)

	// urut data pasien berdasarkan nama
	urutkanNama(A, count)

	// binary search
	low := 0
	high := *count - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if A[mid].nama == pilih {
			fmt.Printf("1. %s %s %d %d\n", A[mid].nama, A[mid].jenis, A[mid].harga, A[mid].tahun)
			found = true
			break
		} else if A[mid].nama < pilih {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Printf("Tidak ada pasien dengan nama %s\n", pilih)
	}
}

//selection sort ascending
func urutkanPasienByTahunAscending(A *tabPasien, count *int) {
	var i, j, idxMin int
	var t pasien
	i = 1
	for i <= *count-1{
		idxMin = i-1
		j = i
		for j < *count{
			if A[idxMin].tahun > A[j].tahun{
				idxMin = j
			}
			j = j+1
		}
		t = A[idxMin]
		A[idxMin] = A[i-1]
		A[i-1] = t
		i = i+1
	}
	fmt.Println()
	fmt.Println("Daftar pasien yang diurutkan berdasarkan tahun (ascending):")
	for i := 0; i < *count; i++ {
		fmt.Print(i+1, ". ")
		fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
	}
	fmt.Println()
}

//selection sort descending
func urutkanPasienByTahunDescending(A *tabPasien, count *int) {
	var i, j, idxMax int
	var t pasien
	i = 1
	for i <= *count-1{
		idxMax = i-1
		j = i
		for j < *count{
			if A[idxMax].tahun < A[j].tahun{
				idxMax = j
			}
			j = j+1
		}
		t = A[idxMax]
		A[idxMax] = A[i-1]
		A[i-1] = t
		i = i+1
	}

	fmt.Println()
	fmt.Println("Daftar pasien yang diurutkan berdasarkan tahun (descending):")
	for i := 0; i < *count; i++ {
		fmt.Print(i+1, ". ")
		fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
	}
	fmt.Println()
}

//insertion sort
func urutkanPasienByPaketAscending(A *tabPasien, count *int) {
	var i, j int
	var temp pasien
	i = 1
	for i <= *count-1{
		j = i
		temp = A[j]
		for j > 0 && temp.harga < A[j-1].harga{
			A[j]=A[j-1]
			j = j-1
		} 
		A[j] = temp
		i = i + 1
	}

	fmt.Println()
	fmt.Println("Daftar pasien yang diurutkan berdasarkan paket (ascending):")
	for i := 0; i < *count; i++ {
		fmt.Print(i+1, ". ")
		fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
	}
	fmt.Println()
}

//insertion sort
func urutkanPasienByPaketDescending(A *tabPasien, count *int) {
	var i, j int
	var temp pasien
	i = 1
	for i <= *count-1{
		j = i
		temp = A[j]
		for j > 0 && temp.harga > A[j-1].harga{
			A[j]=A[j-1]
			j = j-1
		} 
		A[j] = temp
		i = i + 1
	}

	fmt.Println()
	fmt.Println("Daftar pasien yang diurutkan berdasarkan paket (descending):")
	for i := 0; i < *count; i++ {
		fmt.Print(i+1, ". ")
		fmt.Printf("%s %s %d %d\n", A[i].nama, A[i].jenis, A[i].harga, A[i].tahun)
	}
	fmt.Println()
}

func rekap(A *tabPasien, count *int){
	var pilih string
	fmt.Print("Rekap pasien dengan nama: ")
	fmt.Scan(&pilih)
	fmt.Println()
	found := false
	tubuhSehat := "Sehat"
	for i := 0; i < *count; i++ {
		if A[i].nama == pilih {
			fmt.Println("     Nama  : ",A[i].nama)
			fmt.Println("     Paket : ",A[i].jenis)
			fmt.Println("     Harga : ",A[i].harga)
			fmt.Println("     Tahun : ",A[i].tahun)
			fmt.Println()
			fmt.Println("Kondisi Jantung     : ",A[i].jantung)
			fmt.Println("Kondisi Gula Darah  : ",A[i].gulDar)
			fmt.Println("Kondisi Urine       : ",A[i].urine)
			fmt.Println("Kondisi Ginjal      : ",A[i].ginjal)
			fmt.Println("Kondisi Kolestrol   : ",A[i].kolestrol)
			if A[i].jantung == "Normal" && A[i].gulDar == "Normal" && A[i].urine == "Normal" && A[i].ginjal == "Normal" && A[i].kolestrol == "Normal"{
				tubuhSehat = "Sehat"	
				fmt.Println()
				fmt.Println("Kondisi Tubuh       : ",tubuhSehat)
				fmt.Println("Saran               :  Jaga kesehatan dan perbanyak minum air putih.")
				}else if  A[i].jantung == "Tidak_Normal" || A[i].gulDar == "Tidak_Normal" || A[i].urine == "Tidak_Normal" || A[i].ginjal == "Tidak_Normal" || A[i].kolestrol == "Tidak_Normal"{
					tubuhSehat = "Kurang Sehat"
					fmt.Println()
					fmt.Println("Kondisi Tubuh       : ",tubuhSehat)
					fmt.Println()
				}
			if A[i].jantung == "Tidak_Normal"{
				fmt.Println("Saran               :  Periksa ke dokter spesialis jantung")
			}
			if A[i].gulDar == "Tidak_Normal"{
				fmt.Println("Saran Gula Darah    :  Kurangi konsumsi gula per-hari")
			}
			if A[i].urine == "Tidak_Normal"{
				fmt.Println("Saran Urine         :  Perbanyak minum air dan cek ke dokter spesialis urologi")
			}
			if A[i].ginjal == "Tidak_Normal"{
				fmt.Println("Saran Ginjal        :  Periksa ke dokter spesialis ginjal")
			}
			if A[i].kolestrol == "Tidak_Normal"{
				fmt.Println("Saran Kolestrol     :  Kurangi konsumsi lemak jenuh dan periksa ke dokter spesialis")
			}
			found = true
		}
	}
	fmt.Println()
	if !found {
		fmt.Printf("Tidak ada pasien dengan nama %s\n", pilih)
	}
}
