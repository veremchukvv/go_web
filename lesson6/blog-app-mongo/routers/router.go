package routers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_web/lesson5/blog-app/controllers"
	"log"
)

var dsn = "root:1234@tcp(localhost:3306)/blog_app?charset=utf8"

func init() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection OK")
	}

	beego.Router("/", &controllers.BlogController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/post", &controllers.PostController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/post/edit/", &controllers.EditController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/post/new", &controllers.NewController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/post/delete/", &controllers.DeleteController{
		Controller: beego.Controller{},
		Db:         db,
	})

}
