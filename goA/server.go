package main

import (
	"net/http"
	"log"
	// "github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
)

func main() {


	router := NewRouter()

    log.Fatal(http.ListenAndServe(":3754", router))
	// router := echo.New()
	// // e.Use(middleware.Logger())
 	// // e.Use(middleware.Recover())
	// //Router

	// // route := e.POST("/users", func(c echo.Context) error {
	// // 	return nil
	// // })
	// // route.Name = "create-user"
	
	// // // or using the inline syntax
	// // e.GET("/users/:id", func(c echo.Context) error {
	// // 	return nil
	// // }).Name = "get-user"

	// router.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "newtest")
	// })
	// router.GET("/getValue", getValue)
	
	// router.POST("/getValue", getValue2)

	// // e.GET("/api", api)

	// router.GET("/getdate/:a", getdate)

	// router.GET("/divide/:a/:b", divide)
	
	// router.POST("/divide/:a/:b", divide2)


	



	// router.Logger.Fatal(router.Start(":3754"))


	
}
// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	//    }))
	// 	data, err := json.MarshalIndent(e.Routes(), "", "  ")
// if err != nil {
// 	return err
// }
// ioutil.WriteFile("router.json", data, 3754)


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


///เริ่มการใช้ go  https://echo.labstack.com/guide
//  คำสั่ง git https://www.memo8.com/git-basic-command/?fbclid=IwAR2p1s7j2iDw4Jo9soB8JsggJNGOUayxbvzwDLrhvJPRPVkFn7TBkz7eG0E
// go api https://stackoverflow.com/questions/26142074/call-a-function-from-another-package-in-go?fbclid=IwAR3E1l3CnTwVTED_4CU852gwiPKIUmclbRnZLO9xB_vMC8w8WjlJiDT_T28
//ทำเป็นทศนิยม   javascript https://stackoverflow.com/questions/4435170/how-to-parse-float-with-two-decimal-places-in-javascript?fbclid=IwAR2MpMuLscBvvJEInFcny_Rc_gVMHtV2VeP3KJqn9xv7i4Q6o1M8cbyjm7g
