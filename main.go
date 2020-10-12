package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Masukkan lebar dan tinggi gambar : ")
	fmt.Scanf("%d", &num)
	cetakGambar(num)
}

func cetakGambar(param int) {
	if param%2 != 1 || param <= 1 {
		fmt.Println("Lebar dan tinggi harus merupakan bilangan ganjil dan lebih dari 1")
	} else {
		for i := 1; i < param+1; i++ {
			for j := 1; j < param+1; j++ {
				if j == 1 || j == param || i == (param+1)/2 {
					fmt.Print("*  ")
				} else {
					fmt.Print("=  ")
				}
			}
			fmt.Println("\n ")
		}
	}
}
