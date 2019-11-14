package main

import "fmt"

func segitiga_pascal(){
	
	fmt.Println("Membuat Segitiga Asteriks")
	var Angka int
	fmt.Print("Input Baris : ")
	fmt.Scanf("%v", &Angka)

	for i :=0; i<Angka; i++{
		for j := Angka-i; j>=1; j--{
			fmt.Printf(" ")
		}
		for k := 0; k<= i; k++{
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
	
