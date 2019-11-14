package main

import "fmt"

func task1()  {

	fmt.Println("Menghitung luas segitiga tanpa inputan")
	a, t := 10, 20
	L := (a * t) / 2
	fmt.Println(L)
	
	fmt.Println("Menghitung Luas Segitiga")
	var alas, tinggi int 
	fmt.Print("Input alas : ")
	fmt.Scanf("%v", &alas)
	fmt.Print("Input tinggi : ")
	fmt.Scanf("%v", &tinggi)
	Luas_segitiga := (alas * tinggi) / 2
	fmt.Print("Luas segitiga : ")
	fmt.Println(Luas_segitiga)

	fmt.Println("Menghitung luas permukaan tabung tanpa inputan")
	var r, pi float64 = 21, 3.14
	var t1 float64 = 14 
	Lp := 2 * pi * r * (r+ t1)
	fmt.Println(Lp)

	fmt.Println("Menghitung luas permukaan tabung ")
	var jari_jari, tinggi_tabung float64 
	fmt.Print("Input jari-jari : ")
	fmt.Scanf("%v", &jari_jari)
	fmt.Print("Input tinggi tabung : ")
	fmt.Scanf("%v", &tinggi_tabung)

	Luas_permukaan_tabung := 2 * pi * jari_jari * (jari_jari + tinggi_tabung)
	fmt.Print("Luas permukaan tabung : ")
	fmt.Println(Luas_permukaan_tabung)


}
