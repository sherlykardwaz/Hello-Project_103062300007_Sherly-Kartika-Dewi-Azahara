package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
)

type Mahasiswa struct {
	Email     string
	Nama      string
	Jurusan   string
	NoTelepon string
	NilaiTes  float64
	Status    string
}

type Jurusan struct {
	Nama string
}

var mahasiswaList []Mahasiswa
var jurusanList = []Jurusan{
	{"Informatika"},
	{"Teknologi Informasi"},
	{"Sistem Informasi"},
	{"Teknik Telekomunikasi"},
	{"Teknik Elektro"},
	{"Teknik Industri"},
	{"Desain Komunikasi Visual"},
}
var adminPassword = "AllAdmin555"

func main() {
	loadFromTextFile()
	loadFromGobFile()
	defer saveToTextFile()
	defer saveToGobFile()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("   APLIKASI PENDAFTARAN MAHASISWA    ")
		fmt.Println("=====================================")
		fmt.Println("            MENU UTAMA               ")
		fmt.Println("-------------------------------------")
		fmt.Println("1. Daftar Mahasiswa Baru")
		fmt.Println("2. Cek Nilai Tes dan Status Kelulusan")
		fmt.Println("3. Masuk sebagai Admin")
		fmt.Println("4. Exit")
		fmt.Println("=====================================")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			registerMahasiswa(scanner)
		case "2":
			cekNilaiTes(scanner)
		case "3":
			masukAdmin(scanner)
		case "4":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func registerMahasiswa(scanner *bufio.Scanner) {
	var email, nama, jurusan, noTelepon string

	fmt.Print("Email: ")
	scanner.Scan()
	email = scanner.Text()
	if isEmailRegistered(email) {
		fmt.Println("Email sudah terdaftar!")
		return
	}

	fmt.Print("Nama Lengkap: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Println("Pilih Jurusan:")
	for i, jur := range jurusanList {
		fmt.Printf("%d. %s\n", i+1, jur.Nama)
	}
	fmt.Print("Nomor jurusan pilihan: ")
	scanner.Scan()
	jurusanIndex := scanner.Text()
	jurusanInt := toInt(jurusanIndex)
	if jurusanInt < 1 || jurusanInt > len(jurusanList) {
		fmt.Println("Pilihan jurusan tidak valid!")
		return
	}
	jurusan = jurusanList[jurusanInt-1].Nama

	fmt.Print("Nomor telepon: ")
	scanner.Scan()
	noTelepon = scanner.Text()

	mahasiswaList = append(mahasiswaList, Mahasiswa{Email: email, Nama: nama, Jurusan: jurusan, NoTelepon: noTelepon})
	fmt.Println("Berhasil terdaftar!")
}

func cekNilaiTes(scanner *bufio.Scanner) {
	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()
	for _, mhs := range mahasiswaList {
		if mhs.Email == email {
			fmt.Printf("Nilai Tes: %.2f\n", mhs.NilaiTes)
			fmt.Printf("Status Kelulusan: %s\n", mhs.Status)
			return
		}
	}
	fmt.Println("Email tidak ditemukan!")
}

func masukAdmin(scanner *bufio.Scanner) {
	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()
	if password != adminPassword {
		fmt.Println("Password salah!")
		return
	}

	for {
		fmt.Println()
		fmt.Println("==================================================")
		fmt.Println("                    MENU ADMIN                    ")
		fmt.Println("--------------------------------------------------")
		fmt.Println("1. Kelola Data Jurusan")
		fmt.Println("2. Kelola Nilai Tes dan Status Kelulusan Mahasiswa")
		fmt.Println("3. Kelola Data Mahasiswa")
		fmt.Println("4. Tampilkan Data Mahasiswa")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Println("==================================================")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			kelolaJurusan(scanner)
		case "2":
			kelolaNilaiTes(scanner)
		case "3":
			kelolaDataMahasiswa(scanner)
		case "4":
			tampilkanDataMahasiswa(scanner)
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func kelolaJurusan(scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Println("===================================")
		fmt.Println("        KELOLA DATA JURUSAN        ")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Tambah")
		fmt.Println("2. Ubah")
		fmt.Println("3. Hapus")
		fmt.Println("4. Tampilkan Jurusan yang Tersimpan")
		fmt.Println("5. Kembali ke Menu Admin")
		fmt.Println("===================================")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			tambahJurusan(scanner)
		case "2":
			ubahJurusan(scanner)
		case "3":
			hapusJurusan(scanner)
		case "4":
			tampilkanJurusan()
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahJurusan(scanner *bufio.Scanner) {
	fmt.Print("Nama Jurusan: ")
	scanner.Scan()
	nama := scanner.Text()
	jurusanList = append(jurusanList, Jurusan{Nama: nama})
	fmt.Println("Jurusan berhasil ditambahkan!")
}

func ubahJurusan(scanner *bufio.Scanner) {
	fmt.Println("Pilih jurusan yang akan diubah:")
	for i, jur := range jurusanList {
		fmt.Printf("%d. %s\n", i+1, jur.Nama)
	}
	fmt.Print("Nomor jurusan: ")
	scanner.Scan()
	jurusanIndex := scanner.Text()
	jurusanInt := toInt(jurusanIndex)
	if jurusanInt < 1 || jurusanInt > len(jurusanList) {
		fmt.Println("Pilihan jurusan tidak valid!")
		return
	}

	fmt.Print("Nama Jurusan Baru: ")
	scanner.Scan()
	nama := scanner.Text()
	jurusanList[jurusanInt-1].Nama = nama
	fmt.Println("Jurusan berhasil diubah!")
}

func hapusJurusan(scanner *bufio.Scanner) {
	fmt.Println("Pilih jurusan yang akan dihapus:")
	for i, jur := range jurusanList {
		fmt.Printf("%d. %s\n", i+1, jur.Nama)
	}
	fmt.Print("Nomor jurusan: ")
	scanner.Scan()
	jurusanIndex := scanner.Text()
	jurusanInt := toInt(jurusanIndex)
	if jurusanInt < 1 || jurusanInt > len(jurusanList) {
		fmt.Println("Pilihan jurusan tidak valid!")
		return
	}

	jurusanList = append(jurusanList[:jurusanInt-1], jurusanList[jurusanInt:]...)
	fmt.Println("Jurusan berhasil dihapus!")
}

func tampilkanJurusan() {
	fmt.Println("Jurusan yang tersimpan:")
	for i, jur := range jurusanList {
		fmt.Printf("%d. %s\n", i+1, jur.Nama)
	}
}

func kelolaNilaiTes(scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Println("===============================================")
		fmt.Println("KELOLA NILAI TES DAN STATUS KELULUSAN MAHASISWA")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Tambah")
		fmt.Println("2. Ubah")
		fmt.Println("3. Hapus")
		fmt.Println("4. Kembali ke Menu Admin")
		fmt.Println("===============================================")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			tambahNilaiTes(scanner)
		case "2":
			ubahNilaiTes(scanner)
		case "3":
			hapusNilaiTes(scanner)
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahNilaiTes(scanner *bufio.Scanner) {
	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()
	for i, mhs := range mahasiswaList {
		if mhs.Email == email {
			fmt.Print("Nilai Tes: ")
			scanner.Scan()
			nilai := toFloat(scanner.Text())
			fmt.Print("Status Kelulusan (Diterima/Ditolak): ")
			scanner.Scan()
			status := scanner.Text()
			mahasiswaList[i].NilaiTes = nilai
			mahasiswaList[i].Status = status
			fmt.Println("Nilai tes dan status kelulusan berhasil ditambahkan!")
			return
		}
	}
	fmt.Println("Email tidak ditemukan!")
}

func ubahNilaiTes(scanner *bufio.Scanner) {
	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()
	for i, mhs := range mahasiswaList {
		if mhs.Email == email {
			fmt.Print("Nilai Tes Baru: ")
			scanner.Scan()
			nilai := toFloat(scanner.Text())
			fmt.Print("Status Kelulusan Baru: ")
			scanner.Scan()
			status := scanner.Text()
			mahasiswaList[i].NilaiTes = nilai
			mahasiswaList[i].Status = status
			fmt.Println("Nilai tes dan status kelulusan berhasil diubah!")
			return
		}
	}
	fmt.Println("Email tidak ditemukan!")
}

func hapusNilaiTes(scanner *bufio.Scanner) {
	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()
	for i, mhs := range mahasiswaList {
		if mhs.Email == email {
			mahasiswaList[i].NilaiTes = 0
			mahasiswaList[i].Status = ""
			fmt.Println("Nilai tes dan status kelulusan berhasil dihapus!")
			return
		}
	}
	fmt.Println("Email tidak ditemukan!")
}

func kelolaDataMahasiswa(scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Println("========================")
		fmt.Println(" KELOLA DATA MAHASISWA  ")
		fmt.Println("------------------------")
		fmt.Println("1. Edit")
		fmt.Println("2. Hapus")
		fmt.Println("3. Kembali ke Menu Admin")
		fmt.Println("========================")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			editDataMahasiswa(scanner)
		case "2":
			hapusDataMahasiswa(scanner)
		case "3":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func editDataMahasiswa(scanner *bufio.Scanner) {
	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()
	for i, mhs := range mahasiswaList {
		if mhs.Email == email {
			fmt.Print("Nama Lengkap Baru: ")
			scanner.Scan()
			nama := scanner.Text()
			fmt.Println("Pilih Jurusan Baru:")
			for i, jur := range jurusanList {
				fmt.Printf("%d. %s\n", i+1, jur.Nama)
			}
			fmt.Print("Nomor jurusan pilihan: ")
			scanner.Scan()
			jurusanIndex := scanner.Text()
			jurusanInt := toInt(jurusanIndex)
			if jurusanInt < 1 || jurusanInt > len(jurusanList) {
				fmt.Println("Pilihan jurusan tidak valid!")
				return
			}
			jurusan := jurusanList[jurusanInt-1].Nama

			fmt.Print("Nomor telepon baru: ")
			scanner.Scan()
			noTelepon := scanner.Text()

			mahasiswaList[i].Nama = nama
			mahasiswaList[i].Jurusan = jurusan
			mahasiswaList[i].NoTelepon = noTelepon
			fmt.Println("Data mahasiswa berhasil diubah!")
			return
		}
	}
	fmt.Println("Email tidak ditemukan!")
}

func hapusDataMahasiswa(scanner *bufio.Scanner) {
	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()
	for i, mhs := range mahasiswaList {
		if mhs.Email == email {
			mahasiswaList = append(mahasiswaList[:i], mahasiswaList[i+1:]...)
			fmt.Println("Mahasiswa berhasil dihapus!")
			return
		}
	}
	fmt.Println("Email tidak ditemukan!")
}

func tampilkanDataMahasiswa(scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Println("=================================================")
		fmt.Println("             TAMPILAN DATA MAHASISWA             ")
		fmt.Println("-------------------------------------------------")
		fmt.Println("1. Mahasiswa yang mendaftar pada jurusan tertentu")
		fmt.Println("2. Status kelulusan diterima atau ditolak")
		fmt.Println("3. Terurut berdasarkan nilai tes")
		fmt.Println("4. Terurut berdasarkan jurusan")
		fmt.Println("5. Terurut berdasarkan nama")
		fmt.Println("6. Kembali ke Menu Admin")
		fmt.Println("=================================================")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			tampilkanMahasiswaByJurusan(scanner)
		case "2":
			tampilkanMahasiswaByStatus(scanner)
		case "3":
			fmt.Print("Urut secara? (asc/desc): ")
			scanner.Scan()
			order := scanner.Text()
			ascending := true
			if order == "desc" {
				ascending = false
			}
			sortByNilaiTes(ascending)
			tampilkanSemuaMahasiswa()
		case "4":
			fmt.Print("Urut secara? (asc/desc): ")
			scanner.Scan()
			order := scanner.Text()
			ascending := true
			if order == "desc" {
				ascending = false
			}
			sortByJurusan(ascending)
			tampilkanSemuaMahasiswa()
		case "5":
			fmt.Print("Urut secara? (asc/desc): ")
			scanner.Scan()
			order := scanner.Text()
			ascending := true
			if order == "desc" {
				ascending = false
			}
			sortByNama(ascending)
			tampilkanSemuaMahasiswa()
		case "6":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tampilkanMahasiswaByJurusan(scanner *bufio.Scanner) {
	fmt.Println("Pilih Jurusan:")
	for i, jur := range jurusanList {
		fmt.Printf("%d. %s\n", i+1, jur.Nama)
	}
	fmt.Print("Nomor jurusan pilihan: ")
	scanner.Scan()
	jurusanIndex := scanner.Text()
	jurusanInt := toInt(jurusanIndex)
	if jurusanInt < 1 || jurusanInt > len(jurusanList) {
		fmt.Println("Pilihan jurusan tidak valid!")
		return
	}
	jurusan := jurusanList[jurusanInt-1].Nama
	fmt.Printf("\nMahasiswa yang mendaftar pada jurusan %s:\n", jurusan)
	for _, mhs := range mahasiswaList {
		if mhs.Jurusan == jurusan {
			fmt.Printf("Email: %s, Nama: %s, Telepon: %s, Nilai Tes: %.2f, Status: %s\n",
				mhs.Email, mhs.Nama, mhs.NoTelepon, mhs.NilaiTes, mhs.Status)
		}
	}
}

func tampilkanMahasiswaByStatus(scanner *bufio.Scanner) {
	fmt.Print("Status kelulusan (Diterima/Ditolak): ")
	scanner.Scan()
	status := scanner.Text()
	fmt.Printf("\nMahasiswa dengan status kelulusan %s:\n", status)
	for _, mhs := range mahasiswaList {
		if mhs.Status == status {
			fmt.Printf("Email: %s, Nama: %s, Jurusan: %s, Telepon: %s, Nilai Tes: %.2f\n",
				mhs.Email, mhs.Nama, mhs.Jurusan, mhs.NoTelepon, mhs.NilaiTes)
		}
	}
}

func sortByNilaiTes(ascending bool) {
	for i := 0; i < len(mahasiswaList)-1; i++ {
		for j := 0; j < len(mahasiswaList)-i-1; j++ {
			if (ascending && mahasiswaList[j].NilaiTes > mahasiswaList[j+1].NilaiTes) ||
				(!ascending && mahasiswaList[j].NilaiTes < mahasiswaList[j+1].NilaiTes) {
				mahasiswaList[j], mahasiswaList[j+1] = mahasiswaList[j+1], mahasiswaList[j]
			}
		}
	}
}

func sortByJurusan(ascending bool) {
	for i := 0; i < len(mahasiswaList)-1; i++ {
		for j := 0; j < len(mahasiswaList)-i-1; j++ {
			if (ascending && mahasiswaList[j].Jurusan > mahasiswaList[j+1].Jurusan) ||
				(!ascending && mahasiswaList[j].Jurusan < mahasiswaList[j+1].Jurusan) {
				mahasiswaList[j], mahasiswaList[j+1] = mahasiswaList[j+1], mahasiswaList[j]
			}
		}
	}
}

func sortByNama(ascending bool) {
	for i := 0; i < len(mahasiswaList)-1; i++ {
		for j := 0; j < len(mahasiswaList)-i-1; j++ {
			if (ascending && mahasiswaList[j].Nama > mahasiswaList[j+1].Nama) ||
				(!ascending && mahasiswaList[j].Nama < mahasiswaList[j+1].Nama) {
				mahasiswaList[j], mahasiswaList[j+1] = mahasiswaList[j+1], mahasiswaList[j]
			}
		}
	}
}

func tampilkanSemuaMahasiswa() {
	for _, mhs := range mahasiswaList {
		fmt.Printf("Email: %s, Nama: %s, Jurusan: %s, Telepon: %s, Nilai Tes: %.2f, Status: %s\n",
			mhs.Email, mhs.Nama, mhs.Jurusan, mhs.NoTelepon, mhs.NilaiTes, mhs.Status)
	}
}

func saveToTextFile() {
	file, err := os.Create("data_mahasiswa.txt")
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Data Pendaftaran Mahasiswa:\n")
	for _, mhs := range mahasiswaList {
		writer.WriteString(fmt.Sprintf("Email: %s, Nama: %s, Jurusan: %s, Telepon: %s, Nilai Tes: %.2f, Status: %s\n",
			mhs.Email, mhs.Nama, mhs.Jurusan, mhs.NoTelepon, mhs.NilaiTes, mhs.Status))
	}
	writer.Flush()
}

func loadFromTextFile() {
	file, err := os.Open("data_mahasiswa.txt")
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Data Mahasiswa:" {
			continue
		}

		var mhs Mahasiswa
		fmt.Sscanf(line, "Email: %s, Nama: %s, Jurusan: %s, Telepon: %s, Nilai Tes: %f, Status: %s",
			&mhs.Email, &mhs.Nama, &mhs.Jurusan, &mhs.NoTelepon, &mhs.NilaiTes, &mhs.Status)
		mahasiswaList = append(mahasiswaList, mhs)
	}
}

func saveToGobFile() {
	file, err := os.Create("data_mahasiswa.gob")
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(mahasiswaList)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
	}
}

func loadFromGobFile() {
	file, err := os.Open("data_mahasiswa.gob")
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&mahasiswaList)
	if err != nil {
		fmt.Println("Gagal memuat data:", err)
	}
}

func isEmailRegistered(email string) bool {
	for _, mhs := range mahasiswaList {
		if mhs.Email == email {
			return true
		}
	}
	return false
}

func toInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func toFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return f
}