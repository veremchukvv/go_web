package main

import (
	"github.com/astaxie/beego"
	_ "go_web/lesson6/blog-app-mongo/routers"
	"os"
)

func main() { // цикломатическая сложность функции = 1
	beego.Run("localhost:" + os.Getenv("httpport"))
}
