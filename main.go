package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct {
	id        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var biodata = []person{
	{
		id:        1,
		nama:      "Rizky",
		alamat:    "Jl. Raya",
		pekerjaan: "Programmer",
		alasan:    "Mau belajar",
	},
	{
		id:        2,
		nama:      "Ali",
		alamat:    "Jl. Raya Combor no 3",
		pekerjaan: "Freelancer",
		alasan:    "Mau belajar Go Go",
	},
	{
		id:        3,
		nama:      "Budi",
		alamat:    "Jl. Soekarno Hatta",
		pekerjaan: "Programmer",
		alasan:    "Mau belajar Microservice di Go",
	},
	{
		id:        4,
		nama:      "Rudi Habiebie",
		alamat:    "Jl. Rudi Habieboe",
		pekerjaan: "Programmer",
		alasan:    "Mau belajar Go",
	},
}

func main() {

	args := "Data tidak ada"
	if len(os.Args) > 1 {
		args = os.Args[1]
	}

	intArgs, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println("Mohon maaf terjadi kesalahan")
		return
	}

	if intArgs > len(biodata) || intArgs < 1 {
		fmt.Println("Maaf, Data tidak ditemukan")
		return
	}

	for _, v := range biodata {
		if v.id == intArgs {
			fmt.Println(v.nama)
			fmt.Println(v.alamat)
			fmt.Println(v.pekerjaan)
			fmt.Println(v.alasan)
		}
	}
}
