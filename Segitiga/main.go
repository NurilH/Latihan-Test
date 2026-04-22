package main

import "fmt"

func main() {
	SegitigaSamaKaki(4)
	SegitigaSikuKiri(5)
	SegitigaSikuKanan(3)
	SACount()
	SAGanjilGenap()
	SARevers(15)
}

func SegitigaSamaKaki(count int) {
	for i := 1; i <= count; i++ {
		for j := 1; j <= count-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func SegitigaSikuKiri(count int) {
	for i := 0; i < count; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func SegitigaSikuKanan(count int) {
	for i := 0; i < count; i++ {
		for j := 0; j < count-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
