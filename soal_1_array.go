package main

import "fmt"

func soal_1(){
slice := append ([]string {"didin","ayu","sinta","didin", "ayu"},) 
slice_2 := append ([]string {"didin","ayu","sinta","didin", "ayu"}) 

slice_3 := append(slice,slice_2...)
keys := make (map [string]bool)
list := [] string {}

for _, n := range slice_3 { // loop cek slice
	if _, value := keys [n]; !value { // seleksi jika n (slice) tidak sama sama value 
		keys [n]= true  // range size true
		list = append (list,n) 
	}
}
fmt.Println (list)
}