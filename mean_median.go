package main

import "fmt"

func mean_median(){
    var data[100] int
	var banyak_data,sum int
	var avg float64
	
    fmt.Print("Masukkan Banyak Data: ")
	fmt.Scanln(&banyak_data)
	
    for i := 0; i < banyak_data; i++ {
        fmt.Print("Masukkan data : ")
        fmt.Scanln(&data[i])     
        sum += data[i]
    }

	avg = float64(sum)/float64(banyak_data)
	fmt.Printf("Rata-rata dari data yang diinputkan adalah  %v \n",avg)

	if banyak_data % 2 == 0{
		median1 := (banyak_data/2) 
		median2 := (banyak_data/2) - 1
		median := float64(data[median1] + data[median2])/2
		fmt.Printf("Median dari data yang diinputkan adalah %v \n", median )

	}else{
		median := (banyak_data - 1)/2 
		fmt.Printf("Median dari data yang diinputkan adalah %v \n", data[median] )
	}
}