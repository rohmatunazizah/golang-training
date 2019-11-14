package main

import "fmt"

func faktor_bilangan()  {
	// fmt.Printf("\n\n")
	fmt.Println("Menghitung faktor bilangan")
	var Angka int 
	fmt.Print("Input Angka : ")
	fmt.Scanf("%v", &Angka)
	fmt.Println("Faktor dari Angka", Angka)
	for i := 1; i <= Angka; i++ {
		if Angka % i == 0 {
			fmt.Println(i)
		}
	}
}