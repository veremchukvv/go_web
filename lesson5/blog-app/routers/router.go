package routers

import (
	"database/sql"
	"github.com/astaxie/beego"
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

	// router.HandleFunc("/post", s.viewOnePost)
	beego.Router("/post", &controllers.PostController{
		Controller: beego.Controller{},
		Db:         db,
	})

	// router.HandleFunc("/post/edit/{id:[0-9]+}", s.editPost)
	beego.Router("/post/edit/{id:[0-9]+}", &controllers.BlogController{
		Controller: beego.Controller{},
		Db:         db,
	})

	// router.HandleFunc("/post/new", s.newPost)
	beego.Router("/post/edit/new", &controllers.BlogController{
		Controller: beego.Controller{},
		Db:         db,
	})

	// router.HandleFunc("/post/delete", s.deletePost)
	beego.Router("/post/delete/{id:[0-9]+}", &controllers.BlogController{
		Controller: beego.Controller{},
		Db:         db,
	})

}
