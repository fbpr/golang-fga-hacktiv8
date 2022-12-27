package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	nama string
	alamat string
	pekerjaan string
	alasan string
}

func main()  {
	gettingArgs, _ := strconv.Atoi(os.Args[1])

	var people = []Person {
		{nama: "A", alamat: "Jakarta", pekerjaan: "Pelajar", alasan: "Menambah skill"},
		{nama: "B", alamat: "Bandung", pekerjaan: "IT", alasan: "Menambah pengalaman"},
		{nama: "C", alamat: "Tangerang", pekerjaan: "Guru", alasan: "Menambah skill"},
	}

	fmt.Printf("%+v", people[gettingArgs - 1])
}

