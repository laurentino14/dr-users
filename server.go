package main

import (
	"github.com/labstack/echo/v4"
	"github.com/laurentino14/dr-users/routes"
	_ "github.com/lib/pq"
)



func main() {

	E:= echo.New()
	
	E.POST("/create", routes.CreateUser)
        E.GET("/list", routes.ListUsers)

	E.Start(":3131")
}

