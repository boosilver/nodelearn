package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	y := float32(t.Year())
	mon := t.Month()
	var i int = int(mon)
	var g = float32(i)
	d := float32(t.Day())
	sum := y + g + d

	fmt.Print("Input Number A to Divide :")
	var n int
	fmt.Scan(&n)
	var nn = float32(n)
	fmt.Println(nn)
	// fmt.Println("The time is", time.Now())
	fmt.Println("Sum of Date is :", sum)
	f := float32(sum / nn)
	fmt.Println("Welcome to the playground!")
	fmt.Println("Value is :", sum, " / ", n)
	fmt.Printf("Sum is %.4f", f)
}
