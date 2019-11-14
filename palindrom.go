package main

import ( 
    "fmt"
    "strings"
)
func palindrom()  {

	fmt.Println("Palindrom")
	var kata, balikKata string
	
	fmt.Print("Input kata : ")
	fmt.Scanf("%v", &kata)
	kata1 := strings.ToLower(kata)

	for i := len(kata1) - 1; i >= 0; i-- {
		balikKata = balikKata + string(kata1[i])
	} 
		// fmt.Println(balikKata)
	
	if kata1 == balikKata {
		fmt.Println(kata, " is Palindrom")
	} else {
		fmt.Println(kata, "  isn't Palindrom")
	}

}