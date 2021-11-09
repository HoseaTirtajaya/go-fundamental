package main

import (
	"fmt"
)

func main() {
	// var name string // TO DECLARE A VARIABLE WITH EMPTY VALUE
	var name = "Hosea Tirtajaya" // TO DECLARE A RANDOM VALUE WITHIN VARIABLE
	var age = 21

	fmt.Println("Halo saya ", name, ". Saya berumur ", age)

	//DECLARING MULTIPLE VARIABLES
	var (
		firstName = "Hosea"
		lastName  = "Tirtajaya"
	)

	fmt.Println(firstName, lastName)

	//CONSTANT
	const alwaysTrue = true
	fmt.Println(alwaysTrue)

	//const can use const (
	// firstName = hosea
	// ) like var up there to declare multiple constants

	//variable conversion
	var convert = string(firstName[0])
	fmt.Println(convert)

	//type declaration(can be called as alias)
	type noKtp string

	var noKtpHose noKtp = "2002840812340834"
	fmt.Println(noKtpHose)

	//Math Ops
	var a = 10
	var b = 20

	fmt.Println(a * b)

	//ARRAY
	var data [3]int
	data[0] = 123
	data[1] = 435
	data[2] = 543

	fmt.Println(data[0])

	var data2 = [2]string{
		"Hosea",
		"Tirtajaya",
	}

	fmt.Println(data2[0] + " " + data2[1])
	fmt.Println(len(data), len(data2))

	//SLICE
	var months = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	//SLICE MENGUBAH DATA ARRAY ASLI DIATAS.
	var slice1 = months[4:7]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]
	var slice3 = append(slice2, "JANDESMAR")
	fmt.Println(slice3, months)

	//CAPACITY DALAM ARRAY  GABS DITAMBAHIN, JDI KALO RANGE UDA MASUK ke index terakhir, slice yang di append jadi array baru, kalo capacitynya pas, bisa di append.

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Mari"
	newSlice[1] = "Belajar"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice, newSlice)

	//PERBEDAAN DEKLARASI SLICE DAN ARRAY
	//var iniArray = [...]int{1,2,3,4,5} ato var iniArray = [5]int{1,2,3,4,5}
	//var iniSlice = []int(1,2,3,4,5)

	//MAP
	person := map[string]string{
		"name":    "Hosea",
		"address": "Jakarta",
	}

	fmt.Println(person["name"])

	person["title"] = "Programmer"
	fmt.Println(person)

	book := make(map[string]string)

	book["title"] = "THE GREAT WAR"
	book["author"] = "Anonim"
	book["isRead"] = "Y"

	fmt.Println(book)

	delete(book, "isRead")
	fmt.Println(book)

}
