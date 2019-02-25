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
    e.File("/js/highlight.min.js", "js/highlight.min.js")

    e.File("/css/style.css", "css/style.css")
    e.File("/css/an-old-hope.css", "css/an-old-hope.css")
    e.File("/css/github.css", "css/github.css")
 
    e.Start(":1323")
}
