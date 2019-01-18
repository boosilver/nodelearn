package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func controller(c echo.Context) error {
	num := c.Param("num")
	return c.String(http.StatusOK, num)
}

func atime(a int) {

	// var i int
	// fmt.Println("The time is", time.Now())

	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	y := t.Year()
	mon := t.Month()
	var i int = int(mon)
	d := t.Day()
	fmt.Println("Year   :", y)
	fmt.Println("Month  :", mon)
	fmt.Println("Day    :", d)
	var sum = y + i + d
	fmt.Println(sum)
	var suma = sum / a
	fmt.Println(suma)

	// d := time.Now().Date()
	// fmt.Println(d)
	// fmt.Println(t.Format(time.ANSIC))

	// t := time.Date(2015, 02, 21, 23, 10, 52, 211, time.UTC)
	// fmt.Println(t)
	// y := t.Year()
	// mon := t.Month()
	// d := t.Day()
	// h := t.Hour()
	// m := t.Minute()
	// s := t.Second()
	// n := t.Nanosecond()

	// fmt.Println("Year   :", y)
	// fmt.Println("Month  :", mon)
	// fmt.Println("Day    :", d)
	// fmt.Println("Hour   :", h)
	// fmt.Println("Minute :", m)
	// fmt.Println("Second :", s)
	// fmt.Println("Nanosec:", n)
}
func main() {

	var a int
	a = 5
	atime(a)

}
