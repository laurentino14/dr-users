package server

import (
	"github.com/labstack/echo/v4"
	"github.com/laurentino14/dr-users/server/routes"
)

func Server(){
	E:= echo.New()
	
	E.POST("/create", routes.CreateUser)
   	E.GET("/list", routes.ListUsers)

	E.Start(":3131")
	
}