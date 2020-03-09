package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "go_web/lesson5/blog-app/routers"
	"os"
)

func main() {
	beego.Run("localhost" + os.Getenv("httpport"))
}
