package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var a, b, c, d int

func main() {

	for {
		sample()
	}
}
func sample() {
	fmt.Println("========================================================================")
	fmt.Println("Press 1 merge 2 array")
	fmt.Println("Press 2 mean and median")
	fmt.Println("Press 3 csv to map")
	fmt.Println("Press 4 slice with command")
	fmt.Println("Press 5 keluar")
	fmt.Println("========================================================================")
	var inputtinggi int
	fmt.Println("Masukan inputan anda :")
	n, err := fmt.Scanln(&inputtinggi)
	if n < 1 || err != nil {
		fmt.Println("invalid inputtinggi")
		return
	}
	switch inputtinggi {
	case 1:
		mergeArray()
	case 2:
		meanMedian()
		// os.Exit(2)
	case 3:
		csvtomap()
	case 4:
		slicewithcommand()
	case 5:
		os.Exit(2)
	default:
		fmt.Println("pilihan menu tidak ada:")
	}
}

func mergeArray() {
	var inputArray1, inputArray2 int

	fmt.Println("masukkan jumlah array 1 : ")
	fmt.Scanf("%v", &inputArray1)
	fmt.Println("masukkan jumlah array 2 : ")
	fmt.Scanf("%v", &inputArray2)
	//tidak bisa pakai array biasa jika pengisian value input jalan saat runtime
	arrayslice1 := make([]string, inputArray1)
	arrayslice2 := make([]string, inputArray2)
	for i := 0; i < inputArray1; i++ {
		fmt.Println("masukkan data array 1 : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		arrayslice1[i] = scanner.Text()
	}
	for j := 0; j < inputArray2; j++ {
		fmt.Println("masukkan data array 2 : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		arrayslice2[j] = scanner.Text()
	}
	fmt.Println(arrayslice1)
	fmt.Println(arrayslice2)
	gabunganArray := append(arrayslice1, arrayslice2...)
	fmt.Println(gabunganArray)
	for k := 0; k < len(gabunganArray); k++ {
		for l := k + 1; l < len(gabunganArray); l++ {
			if gabunganArray[k] == gabunganArray[l] {
				gabunganArray[l] = gabunganArray[len(gabunganArray)-1]
				gabunganArray[len(gabunganArray)-1] = ""
				gabunganArray = gabunganArray[:len(gabunganArray)-1]
			}
		}
	}
	fmt.Println("hasi akhir", gabunganArray)
}

func meanMedian() {
	var inputan1 int
	var median, mean float32

	fmt.Println("masukkan jumlah array 1 : ")
	fmt.Scanf("%v", &inputan1)
	arrayslice1 := make([]float32, inputan1)
	for i := 0; i < inputan1; i++ {
		fmt.Println("masukkan data angka: ")
		fmt.Scanf("%v", &arrayslice1[i])
		mean += arrayslice1[i] / float32(inputan1)
	}
	if inputan1%2 == 0 {
		hasilbagi := inputan1 / 2
		median = (arrayslice1[hasilbagi] + arrayslice1[hasilbagi-1]) / 2
	} else {
		hasilbagi := math.Floor(float64(inputan1) / 2.0)
		median = arrayslice1[int64(hasilbagi)]
	}
	fmt.Printf("hasil median %v: \n ", median)
	fmt.Printf("hasil mean %.2f: \n ", mean)
}

func csvtomap() {
	var inputkeys, inputvalue string
	var inputkeys1 []string
	var inputvalue1 []string
	var mapbaru = map[string]string{}

	fmt.Println("masukkan keys : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputkeys = scanner.Text()
	fmt.Println("masukkan value : ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	inputvalue = scanner1.Text()
	inputkeys1 = strings.Split(inputkeys, ",")
	inputvalue1 = strings.Split(inputvalue, ",")
	for i := 0; i < len(inputkeys1); i++ {
		mapbaru[inputkeys1[i]] = inputvalue1[i]
	}
	fmt.Println(mapbaru)
}

func slicewithcommand() {
	var inputanCommand string
	var dataArray1 []string
	var dataArray2 []string
	dataArray3 := make([]int, 3, 3)

	fmt.Println("masukkan command string : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputanCommand = scanner.Text()
	dataArray1 = strings.Split(inputanCommand, ",")
	fmt.Println(len(dataArray1))
	for i := 0; i < len(dataArray1); i++ {
		dataArray2 = strings.Split(strings.Trim(dataArray1[i], " "), " ")
		switch dataArray2[0] {
		case "insert":
			posisi, _ := strconv.Atoi(dataArray2[1])
			value, _ := strconv.Atoi(dataArray2[2])
			dataArray3[posisi-1] = value
			fmt.Println(dataArray3)
		case "remove":
			posisi, _ := strconv.Atoi(dataArray2[1])
			dataArray3[posisi-1] = dataArray3[len(dataArray3)-1]
			dataArray3[len(dataArray3)-1] = 0
			dataArray3 = dataArray3[:len(dataArray3)-1]
			// jumlah := len(dataArray3) - 1
			// dataArray3 = dataArray3[:jumlah]
			fmt.Println(dataArray3)
		case "append":
			value, _ := strconv.Atoi(dataArray2[1])
			fmt.Println(value)
			dataArray3 = append(dataArray3, value)
			fmt.Println(dataArray3)
		case "sort":
			sort.Ints(dataArray3)
			fmt.Println(dataArray3)
		case "pop":
			dataArray3[len(dataArray3)-1] = 0
			dataArray3 = dataArray3[:len(dataArray3)-1]
			fmt.Println(dataArray3)
		case "reverse":
			sort.Ints(dataArray3)
			sort.Slice(dataArray3, func(i, j int) bool {
				return dataArray3[i] > dataArray3[j]
			})
			fmt.Println(dataArray3)
		}
	}
}
