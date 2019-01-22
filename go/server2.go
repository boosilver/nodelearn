package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func getdate(c echo.Context) error {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	y := float64(t.Year())
	mon := t.Month()
	var l int = int(mon)
	k := float64(l)
	d := float64(t.Day())
	date := fmt.Sprintf("", y, k, d)
	return c.String(http.StatusOK, date)
}

func getsum(c echo.Context) error {
	// User ID from path `users/:id`
	a := c.Param("a")
	g, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	i := float64(g)
	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	y := float64(t.Year())
	mon := t.Month()
	var l int = int(mon)
	k := float64(l)
	d := float64(t.Day())
	fmt.Println("Year   :", y)
	fmt.Println("Month  :", mon)
	fmt.Println("Day    :", d)
	var sum float64 = y + k + d
	fmt.Println("Sum of Date = ", sum)
	suma2 := float64(sum / i)

	s := fmt.Sprintf("%.4f", suma2)
	return c.String(http.StatusOK, s)
}

func divide(c echo.Context) error {
	a := c.Param("a")
	b := c.Param("b")
	f, err := strconv.ParseFloat(a, 64)
	if err == nil {
		fmt.Println("a =", f)
	}
	g, err := strconv.ParseFloat(b, 64)
	if err == nil {
		fmt.Println("b =", g)
	}

	var sum = f / g
	s := fmt.Sprintf("%.4f", sum)
	return c.String(http.StatusOK, s)
}

func main() {
	e := echo.New()
	e.GET("/date", getdate)
	e.GET("/date/:a", getsum)
	e.GET("/divide/:a/:b", divide)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello This is HomePage")
	})
	e.Logger.Fatal(e.Start(":1234"))

}
