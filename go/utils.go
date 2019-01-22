package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// e.GET("/date/:a", getdate)
// e.GET("/divide/:a/:b", divide)

func getValue(c echo.Context) error {
	return c.String(http.StatusOK, "Hello ")

}
func getValue2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello2 ")

}



	func date(c echo.Context) error {
	

		t := time.Now()
		fmt.Println(t.Format("2006-01-02"))
		d :=t.Format("2006-01-02")
		
		return c.String(http.StatusOK, d)
	}
	func postdate(a float64)float64{

		t := time.Now()
		fmt.Println(t.Format("2006-01-02"))
		yy := float64(t.Year())
		mon := t.Month()
		
		mm := float64(t.Month())
		dd := float64(t.Day())
		fmt.Println("Year   :", yy)
		fmt.Println("Month  :", mon)
		fmt.Println("Day    :", dd)
		var sum float64 = yy + mm + dd
		fmt.Println("Year+Month+Day =\t", sum)
		suma2 := float64(sum / a)
		fmt.Print(sum, "/", a)
		fmt.Printf("  result is\t%.3f", suma2)
	
		return suma2
		}

func divide(c echo.Context) error {

	a := c.Param("a")
	b := c.Param("b")
	// d, err := strconv.Atoi(a)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// e, err := strconv.Atoi(b)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// f := float64(d)
	// g := float64(e)

	f, err := strconv.ParseFloat(a, 64)
	if err == nil {  /// ถ้าไม่มีค่าเออเร่อ ให้ทำ
		fmt.Println("a =", f)
	}
	g, err := strconv.ParseFloat(b, 64)
	if err == nil {
		fmt.Println("b =", g)
	}

	var sum = f / g
	s := fmt.Sprintf("%.2f", sum)
	return c.String(http.StatusOK, s)
}

func getdate(c echo.Context) error {
	a := c.Param("a")
	// g, err := strconv.Atoi(a)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// i := float64(g)

	i, err := strconv.ParseFloat(a, 64)
	if err == nil {
		fmt.Println("a =", i)
	}

	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	y := float64(t.Year())
	mon := t.Month()
	// var l int = int(mon)
	k := float64(t.Month())
	d := float64(t.Day())
	fmt.Println("Year   :", y)
	fmt.Println("Month  :", mon)
	fmt.Println("Day    :", d)
	var sum float64 = y + k + d
	fmt.Println("Year+Month+Day =\t", sum)
	suma2 := float64(sum / i)
	fmt.Print(sum, "/", i)
	fmt.Printf("  result is\t%.3f", suma2)

	// t := strconv.Itoa(suma2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	s := fmt.Sprintf("%.2f", suma2)
	return c.String(http.StatusOK, s)
}

type Message struct {
	Text string `json:"text"`
	
}
type Message2 struct {
	Text float64 `json:"text"` 
	
}
func power(c echo.Context) error {

	
	a := c.Param("a")
	b := c.Param("b")
	// d, err := strconv.Atoi(a)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// e, err := strconv.Atoi(b)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// f := float64(d)
	// g := float64(e)

	f, err := strconv.ParseFloat(a, 64)
	if err == nil {  /// ถ้าไม่มีค่าเออเร่อ ให้ทำ
		fmt.Println("a =", f)
	}
	g, err := strconv.ParseFloat(b, 64)
	if err == nil {
		fmt.Println("b =", g)
	}
	var sum float64 =1
	var o float64 =1
	for o=g;o>0;o--{
		// sum :=multiply(sum,f)
		// fmt.Println( sum)
			sum*=f
			
	}
	fmt.Println( f,"^",g,"=",sum)
	// var sum = f / g
	s := fmt.Sprintf("%.2f", sum)
	return c.String(http.StatusOK, s)
}

// func multiply(sum float64 ,f float64)string{
// sum*=f
// newsum := fmt.Sprintf("%.2f", sum)
// return newsum
// }