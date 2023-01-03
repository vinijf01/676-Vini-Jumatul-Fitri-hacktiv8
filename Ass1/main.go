package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	data := []Biodata{
		{Nama: "Ataline Jeanethe Maya Hukubun", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Ade Putri Mirab Suni", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Akmal Dira Thursina", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Alda Nur Firdaus Ambarwati", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Alfi Nurul Izzah", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Alfiyatuz Zamhariroh", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Alif Abdul Bari", Alamat: "Jogja", Pekerjaan: "Freelancer", Alasan: "menambah ilmu dan untuk mencari kerja"},
		{Nama: "Alima Fahmi Rahmawati", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "menambah pengetahuan dalam service suatu website"},
		{Nama: "Alyad Ulya Iman", Alamat: "Bogor", Pekerjaan: "Freelance", Alasan: "untuk menambah pengetahuan backend"},
		{Nama: "Amelia Sari", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "Menambah pengetahuan dan pemahaman bahasa Go"},
		{Nama: "Ani Susanti", Alamat: "Pamulang", Pekerjaan: "Freelance", Alasan: "untuk menambah pengetahuan backend"},
		{Nama: "Anne Sagitariyanti ", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "Menambahskill di bidang IT"},
		{Nama: "Annisa Nuri Nabila", Alamat: "Padang", Pekerjaan: "Freelance", Alasan: "ingin mengenal pemrograman golang "},
		{Nama: "Arnold Julian", Alamat: "Padang", Pekerjaan: "Freelance", Alasan: "ingin mengenal pemrograman golang "},
		{Nama: "Vini Jumatul Fitri", Alamat: "Padang", Pekerjaan: "Freelance", Alasan: "menambah bahasa pemrograman"},
	}

	in := os.Args
	input, _ := strconv.Atoi(in[1])

	fmt.Println("Nama: ", data[input].Nama)
	fmt.Println("Alamat: ", data[input].Alamat)
	fmt.Println("Pekerjaan: ", data[input].Pekerjaan)
	fmt.Println("Alasan Mengikuti DTS: ", data[input].Alasan)
}
