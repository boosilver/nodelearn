package main

import (
	"net/http"
	"github.com/labstack/echo"
	"time"
	"strconv"
	"fmt"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World! i am go ")
	})
	e.GET("/serverdate", getServerTime)
	//e.POST("/sumsum",AmodSUM)
	e.GET("/sumsum/:a",gettypetime)
	e.Logger.Fatal(e.Start(":1234"))
	//e.GET("/testtype", testtype) //int to string ได้แล้วววววววววววววววววว
	//e.GET("/testtype", testtype)  //string to int ได้สักทีโว้ยยยย
}

func getServerTime(c echo.Context) error {
	t := time.Now().Format("2006-01-02 03:04:05")
	return c.String(http.StatusOK, "Time now : " + t)
}
//func testtype(c echo.Context) error {
	//var i int = 12345 ;
	//strInt := strconv.Itoa(i)
	//return c.String(http.StatusOK, "Time now : " + strInt)   //int เป็น string
//}

//func testtype(c echo.Context) error {
	//str  :=  "1";
	//intI, err := strconv.Atoi (str)
	//if err == nil{
	//	fmt.Println(intI)
	//}else{
	//	fmt.Println(err)
	//}
	//intI = intI+5
	//fmt.Println(intI)
	//strInt := strconv.Itoa(intI)
	//return c.String(http.StatusOK, "Time now : " + strInt)   // string to int ได้แล้ววว
//}


func gettypetime(c echo.Context) error {
	var yy,mm,dd,aa = 0.0,0.0,0.0,0.0 ;

	a := c.Param("a")
	 aa, err := strconv.ParseFloat(a, 64)
	if err == nil {
		fmt.Println(aa) }

	y := time.Now().Format("2006")
	 yy, err = strconv.ParseFloat(y, 64)
	if err == nil {
		fmt.Println(yy) }

	m := time.Now().Format("01")
	mm, err = strconv.ParseFloat(m, 64) 
	if err == nil {
		fmt.Println(mm) }

	d := time.Now().Format("02")
	dd, err = strconv.ParseFloat(d, 64) 
	if err == nil {
		fmt.Println(dd) }
	sum := yy+mm+dd;
	suma := sum/aa;
	r := fmt.Sprintf("%.4f", suma)

	return c.String(http.StatusOK, "sumsum : " + r) 




	
	//yy:= strconv.ParseInt(y)
	//mm:= strconv.ParseInt(m)
	//yy, err := strconv.Atoi(y)
	//a := yy+mm
	
	//return y,m,d
	//strInt := strconv.Itoa(err)
	//return c.String(http.StatusOK,"a is "+a )
 }



// func AmodSUM(c echo.Context) error {

	//s := fmt.Sprintf("%.2f", 12.3456) // s == "12.35" แปลงทศนิยม


// }