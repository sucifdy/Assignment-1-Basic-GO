package main

import (
	"fmt"
	"strings"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	for _, student := range strings.Split(Students, ", ") {
		parts := strings.Split(student, "_")
		if parts[0] == id && parts[1] == name {
			return fmt.Sprintf("Login berhasil: %s (%s)", name, parts[2])
		}
	}
	return "Login gagal: data mahasiswa tidak ditemukan"
}

func Register(id string, name string, major string) string {
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	for _, student := range strings.Split(Students, ", ") {
		parts := strings.Split(student, "_")
		if parts[0] == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	Students += fmt.Sprintf(", %s_%s_%s", id, name, major)
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {
	if code == "" {
		return "Code is undefined!"
	}

	for _, program := range strings.Split(StudentStudyPrograms, ", ") {
		parts := strings.Split(program, "_")
		if parts[0] == code {
			return parts[1]
		}
	}
	return "Program study tidak ditemukan"
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))
		case "2":
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))
		case "3":
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
