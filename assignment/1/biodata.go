package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Name      string
	Alamat    string
	Pekerjaan string
}

type Person struct {
	person Biodata
	Alasan string
}

func main() {
	var check_input = os.Args[1]

	if check_input == "1" {
		firstPerson()
	} else if check_input == "2" {
		secondPerson()
	} else if check_input == "3" {
		thirdPerson()
	} else {
		dataNotFound()
	}

}

func firstPerson() {
	var firstPerson Person

	firstPerson.person.Name = "Kresna"
	firstPerson.person.Alamat = "Ilir Barat I Palembang"
	firstPerson.person.Pekerjaan = "Software Engineer"
	firstPerson.Alasan = "Ingin memperdalam bahasa Go"

	fmt.Println("Nama :", firstPerson.person.Name)
	fmt.Println("Alamat :", firstPerson.person.Alamat)
	fmt.Println("Pekerjaan :", firstPerson.person.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang :", firstPerson.Alasan)
}

func secondPerson() {
	var secondPerson Person

	secondPerson.person.Name = "Vespri"
	secondPerson.person.Alamat = "Kemuning Palembang"
	secondPerson.person.Pekerjaan = "Teaching Asistant"
	secondPerson.Alasan = "Ingin mengetahui bahasa Go lebih lanjut"

	fmt.Println("Nama :", secondPerson.person.Name)
	fmt.Println("Alamat :", secondPerson.person.Alamat)
	fmt.Println("Pekerjaan :", secondPerson.person.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang :", secondPerson.Alasan)
}

func thirdPerson() {
	var thirdPerson Person

	thirdPerson.person.Name = "Wijaya"
	thirdPerson.person.Alamat = "Bukit Besar Palembang"
	thirdPerson.person.Pekerjaan = "Full Stack Engineer"
	thirdPerson.Alasan = "Menguasai basic, algoritma, OOP dan seluruh hal tentang Go language"

	fmt.Println("Nama :", thirdPerson.person.Name)
	fmt.Println("Alamat :", thirdPerson.person.Alamat)
	fmt.Println("Pekerjaan :", thirdPerson.person.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang :", thirdPerson.Alasan)
}

func dataNotFound() {
	fmt.Println("Harap masukkan 1, 2, atau 3 setelah nama file !!!")
}
