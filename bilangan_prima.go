package main

import "fmt"

func bilangan_prima()  {
	fmt.Println("Menghitung bilangan prima")
	var Angka int 
	jumlah_faktor := 0
	fmt.Print("Input Angka : ")
	fmt.Scanf("%v", &Angka)

	for i := 1; i <= Angka; i++ {
		if Angka <= 1{
			fmt.Println("Bukan Bilangan Prima")
		} else if Angka % i == 0 {
			jumlah_faktor++
		}
	}
	if jumlah_faktor == 2 {
		fmt.Println("Bilangan Prima")
	} else{
		fmt.Println("Bukan Bilangan Prima")
	}
	
}