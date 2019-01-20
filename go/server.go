package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "newtest")
	})
	e.GET("/getValue", getValue)
	e.GET("/api", api)
	e.GET("/getdate/:a", getdate)
	e.GET("/divide/:a/:b", divide)

	e.Logger.Fatal(e.Start(":3754"))

}

//ต้องโยนฟังก์ชั่นไป

// func controller(c echo.Context) error {
// 	num := c.Param("num")
// 	return c.String(http.StatusOK, num)
// }
// router.New()
// router := mux.NewRouter()
// router.HandleFunc("/api", controller).Methods("GET")
// getdate("/date/:a", getdate)
// e = router
// e.getValue()
// zz := router.getValue()
// fmt.Println(zz)
// router.getValue()
// e.GET("/divide/:a", router.getdate())

// e.GET("/getdate/:num", getdate)

// e.GET("/controller2", func(c echo.Context) error {
// 	return c.String(http.StatusOK, "newkuyboss2")
// })

// router.GET("/", Index)
// router.GET("/hello/:name", Hello)
// log.Fatal(http.ListenAndServe(":3754", router))
