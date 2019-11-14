package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// var a, b, c, d int

// func main() {

// 	for {
// 		sample()
// 	}
// }
// func sample() {
// 	fmt.Println("========================================================================")
// 	fmt.Println("Press 1 mencaari bilangan prima")
// 	fmt.Println("Press 2 mix and max")
// 	fmt.Println("Press 3 romawi")
// 	fmt.Println("Press 4 slice with command")
// 	fmt.Println("Press 5 keluar")
// 	fmt.Println("========================================================================")
// 	var inputtinggi int
// 	fmt.Println("Masukan inputan anda :")
// 	n, err := fmt.Scanln(&inputtinggi)
// 	if n < 1 || err != nil {
// 		fmt.Println("invalid inputtinggi")
// 		return
// 	}
// 	switch inputtinggi {
// 	case 1:
// 		getPrime()
// 	case 2:
// 		mixandmax()
// 		// os.Exit(2)
// 	case 3:
// 		romawi()
// 	case 4:
// 		os.Exit(2)
// 	default:
// 		fmt.Println("pilihan menu tidak ada:")
// 	}
// }

// func getPrime() {
// 	var j = 2
// 	var counter1 = 0
// 	var inputan int
// 	var databaru int

// 	fmt.Println("masukkan index bilangan prima : ")
// 	fmt.Scanf("%v", &inputan)
// 	for {
// 		var counter = 0
// 		for i := 2; i <= j; i++ {
// 			if j%i == 0 {
// 				counter++
// 			}
// 		}
// 		if inputan == 1 {
// 			databaru = 2
// 			fmt.Printf("data prima ke %v adalah %v \n", inputan, databaru)
// 			os.Exit(2)
// 		} else if counter == 1 {
// 			counter1++
// 			databaru = j
// 		} else if counter1 == inputan {
// 			fmt.Printf("data prima ke %v adalah %v \n", counter1, databaru)
// 			os.Exit(2)
// 		}
// 		j++
// 	}

// }

// func mixandmax() {
// 	var inputanAngka int
// 	fmt.Println("Masukkan inputan angka :")
// 	fmt.Scanf("%v", &inputanAngka)
// 	arrayData := make([]int, inputanAngka)
// 	for i := 0; i < inputanAngka; i++ {
// 		fmt.Println("Masukkan inputan angka ke-", (i + 1))
// 		fmt.Scanf("%v", &arrayData[i])
// 	}
// 	fmt.Println(arrayData)
// 	maxmin := func(data []int) (int, int) {
// 		sort.Ints(data)
// 		return data[0], data[1]
// 	}
// 	max, min := maxmin(arrayData)
// 	fmt.Println("nilai minimumnya : ", min)
// 	fmt.Println("nilai maximumnya : ", max)
// }

// func romawi() {
// 	var inputangka int
// 	var romawiawal = ""
// 	fmt.Println("masukkan angka : ")
// 	fmt.Scanf("%v", &inputangka)
// 	defer func() {
// 		fmt.Println("merupakan angka romawinya ")
// 	}()
// 	desimal := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
// 	romawi := []string{"M", "CM", "D", "CD", "X", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
// 	koversiromawi := func(inputangka int) (string, error) {
// 		if inputangka > 1000 || inputangka < 1 {
// 			return "-1", errors.New("number should not be less than 1000 or more than 1")
// 		}
// 		for i := 0; i < len(desimal); i++ {
// 			for desimal[i] <= inputangka {
// 				romawiawal += romawi[i]
// 				inputangka -= desimal[i]
// 			}
// 		}
// 		return romawiawal, nil
// 	}
// 	result, err := koversiromawi(inputangka)
// 	if err == nil {
// 		fmt.Print("ini : ", result, " ")
// 	} else {
// 		fmt.Print("error : ", err)
// 	}

// }
