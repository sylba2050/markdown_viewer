package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.File("/", "html/index.html")

    e.File("/js/main.js", "js/main.js")

    e.File("/css/style.css", "css/style.css")
    e.File("/css/code.css", "css/code.css")
    e.File("/css/github.css", "css/github.css")
 
    e.Start(":1323")
}
