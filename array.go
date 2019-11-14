package main

import "fmt"

func array()  {
	fmt.Println("==============================")
	a := []string{"senin","selasa","rabu","kamis","jumat"}
	b := []string{"senin","selasa","sabtu", "minggu"}

	for i:= 0; i<len(a); i++{
		for j:=0; j<len(b); j++{
			if string(a[i]) == string(b[j]){
				// c:= a[i] 
				// fmt.Printf("data %v",c)
			}else {
				// c:= append(c,"")
				d := b[j]
				fmt.Println(d[j])
			}
		}
	}
	// fmt.Println("elements : %v ", c)
}