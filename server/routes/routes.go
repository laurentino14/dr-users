package routes

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)



type User struct {
	Nome string `json:"nome"`
	Id int `json:"id"`

}


	func CreateUser(c echo.Context) error{
		user := User{}
		c.Bind(&user)
		err := persistUser(user)
		if err != nil{
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusCreated, user)
	}
	
	
	func ListUsers(c echo.Context) error{
		db, err := sql.Open("postgres","host=localhost port=5432 user=postgres "+
		"password=dev dbname=users sslmode=disable")
		if err != nil{
			return err
		}
	
		var resultado []User
		var res User
	
		defer db.Close()
	
		query, err := db.Query("SELECT * FROM users ORDER BY (id);")
		if err != nil{
			return c.JSON(http.StatusNotFound,err)
		}
		for query.Next(){
			query.Scan(&res.Nome,&res.Id)
			resultado = append(resultado,res)
		}
		return c.JSONPretty(200,resultado," ")
	}
	
	
	
	func persistUser(u User) error {
		db, err := sql.Open("postgres","host=localhost port=5432 user=postgres "+
		"password=dev dbname=users sslmode=disable")
		if err != nil{
			return err
		}
	
		defer db.Close()
	
		
		stmt, err := db.Prepare("INSERT INTO users (nome,id) VALUES ($1, $2)")
		if err != nil{
			return err
		}
		_,err = stmt.Exec(u.Nome,u.Id)
		if err != nil{
			return err
		}
		return nil
		
	
	}