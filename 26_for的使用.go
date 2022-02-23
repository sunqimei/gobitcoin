package main

import "fmt"

func main() {
	fmt.Println("Hello")

	temp:=0
	for i:=0;i<=100;i++ {
		temp=temp+i
	}
	fmt.Printf("temp的值是：%T", 'a')
}

