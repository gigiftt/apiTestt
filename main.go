package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type User struct {
	FirstName string `json:"FirstName" db:"FirstName"`
    LastName string `json:"LastName" db:"LastName"`
	Age  int    `json:"Age" db:"Age"`
    Email string `json:"Email" db:"Email"`
}

const(
    addUser = "INSERT INTO user (FirstName, LastName, Age, Email) VALUES (? , ? , ? , ?)"
)

func main() {
    e := echo.New()

    db, err := sqlx.Connect("mysql", "root:@(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatalln(err)
	}

    e.GET("/", func(c echo.Context) error {
    
        users := []User{}
        // e.Logger.Fatal("pass")
        err := db.Select(&users, "Select * From user")
        if err != nil {
            log.Fatalln(err)
            return nil
        }
    
        return c.JSON(http.StatusOK, users)
    })

    // e.GET("/")

    e.POST("/", func(c echo.Context) error  {
        users := new(User)

        err := c.Bind(&users); if err != nil {
            return c.String(http.StatusBadRequest, "bad request")
        }

        log.Print(users)

        //func ไหนมี err เพิ่ม log err + return err ด้วย
        db.Exec(addUser,users.FirstName,users.LastName,users.Age,users.Email)

        return c.JSON(http.StatusOK, "success")
    })


    e.GET("/filter/:name", func(c echo.Context) error {
        users := []User{}

        log.Print(c.Param("name"))
         n := c.Param("name")

        err := db.Select(&users, "Select * From user Where FirstName = ?",n)
        if err != nil {
            log.Fatalln(err)
            return nil
        }

        // result,_ := db.Query("Select FirstName, LastName, Age, Email From user Where FirstName = 'Sam'")

        // log.Print(result.Columns())

        return c.JSON(http.StatusOK, users)
    })


    e.POST("/filter/:name", func(c echo.Context) error {
        users := new(User)

        err := c.Bind(&users); if err != nil {
            return c.String(http.StatusBadRequest, "bad request")
        }

        db.Exec("Update user SET FirstName = ?, LastName = ?, Age = ?, Email=? where LastName = ?",users.FirstName,users.LastName,users.Age,users.Email, c.Param("name"))
        if err != nil {
            log.Fatalln(err)
            return nil
        }

        return c.JSON(http.StatusOK, "success")
    })

    


    e.Logger.Fatal(e.Start(":8080"))
}

